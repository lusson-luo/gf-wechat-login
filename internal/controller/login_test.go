package controller_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "login-demo/api/v1"
	"login-demo/internal/controller"
	"login-demo/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"

	_ "login-demo/internal/packed"
)

func init() {
	fmt.Println("========= init test.login")
}

var (
	db         gdb.DB
	dbPrefix   gdb.DB
	dbInvalid  gdb.DB
	configNode gdb.ConfigNode
	dbDir      = gfile.Temp("sqlite")
	ctx        = gctx.New()

	// Error
	ErrorSave = gerror.NewCode(gcode.CodeNotSupported, `Save operation is not supported by sqlite driver`)
)

const (
	TableSize               = 10
	TableName               = "user"
	TableNameWhichIsKeyword = "group"
	TestSchema1             = "test1"
	TestSchema2             = "test2"
	TableNamePrefix         = "gf_"
	CreateTime              = "2018-10-24 10:00:00"
	DBGroupTest             = "test"
	DBGroupPrefix           = "prefix"
	DBGroupInvalid          = "invalid"
	InitSql                 = "../../manifest/db/db.sql"
)

func init() {
	fmt.Println("init sqlite db start")

	if err := gfile.Mkdir(dbDir); err != nil {
		gtest.Error(err)
	}

	fmt.Println("init sqlite db dir: ", dbDir)

	dbFilePath := gfile.Join(dbDir, "test.db")
	configNode = gdb.ConfigNode{
		Type:    "mysql",
		Link:    fmt.Sprintf(`sqlite::@file(%s)`, dbFilePath),
		Charset: "utf8",
	}
	nodePrefix := configNode
	nodePrefix.Prefix = TableNamePrefix

	nodeInvalid := configNode

	gdb.AddConfigNode(DBGroupTest, configNode)
	gdb.AddConfigNode(DBGroupPrefix, nodePrefix)
	gdb.AddConfigNode(DBGroupInvalid, nodeInvalid)
	gdb.AddConfigNode(gdb.DefaultGroupName, configNode)

	// Default db.
	if r, err := gdb.NewByGroup(); err != nil {
		gtest.Error(err)
	} else {
		db = r
	}

	// Prefix db.
	if r, err := gdb.NewByGroup(DBGroupPrefix); err != nil {
		gtest.Error(err)
	} else {
		dbPrefix = r
	}

	// Invalid db.
	if r, err := gdb.NewByGroup(DBGroupInvalid); err != nil {
		gtest.Error(err)
	} else {
		dbInvalid = r
	}

	fmt.Println("init sqlite db finish")
}

func createTableFromSql() {
	// 初始化数据脚本
	sqlFile, err := os.Open(InitSql)
	if err != nil {
		panic(err)
	}
	defer sqlFile.Close()

	sqlBytes, err := ioutil.ReadAll(sqlFile)
	if err != nil {
		panic(err)
	}

	// 执行 SQL 脚本
	_, err = db.Exec(ctx, string(sqlBytes))
	if err != nil {
		panic(err)
	}
}

func createTable(table ...string) string {
	return createTableWithDb(db, table...)
}

func createInitTable(table ...string) string {
	return createInitTableWithDb(db, table...)
}

func dropTable(table string) {
	dropTableWithDb(db, table)
}

func createTableWithDb(db gdb.DB, table ...string) (name string) {
	if len(table) > 0 {
		name = table[0]
	} else {
		name = fmt.Sprintf(`%s_%d`, TableName, gtime.TimestampNano())
	}
	dropTableWithDb(db, name)

	if _, err := db.Exec(ctx, fmt.Sprintf(`
	CREATE TABLE %s (
		id          INTEGER       PRIMARY KEY AUTOINCREMENT
									UNIQUE
									NOT NULL,
		passport    VARCHAR(45)  NOT NULL
									DEFAULT passport,
		password    VARCHAR(128) NOT NULL
									DEFAULT password,
		nickname    VARCHAR(45),
		create_time DATETIME
	);
	`, db.GetCore().QuoteWord(name),
	)); err != nil {
		gtest.Fatal(err)
	}

	return
}

func createInitTableWithDb(db gdb.DB, table ...string) (name string) {
	name = createTableWithDb(db, table...)
	array := garray.New(true)
	for i := 1; i <= TableSize; i++ {
		array.Append(g.Map{
			"id":          i,
			"passport":    fmt.Sprintf(`user_%d`, i),
			"password":    fmt.Sprintf(`pass_%d`, i),
			"nickname":    fmt.Sprintf(`name_%d`, i),
			"create_time": gtime.NewFromStr(CreateTime).String(),
		})
	}

	result, err := db.Insert(ctx, name, array.Slice())
	gtest.AssertNil(err)

	n, e := result.RowsAffected()
	gtest.Assert(e, nil)
	gtest.Assert(n, TableSize)
	return
}

func dropTableWithDb(db gdb.DB, table string) {
	if _, err := db.Exec(ctx, fmt.Sprintf("DROP TABLE IF EXISTS `%s`", table)); err != nil {
		gtest.Error(err)
	}
}

func dropAllTableWithDb() {
	tables, err := db.Tables(ctx)
	if err != nil {
		gtest.Error(err)
	}
	fmt.Println(tables)
	for _, table := range tables {
		// SQLite 中的特殊表，用于存储自增列的当前值，无法删除
		if table == "sqlite_sequence" {
			continue
		}
		if _, err := db.Exec(ctx, fmt.Sprintf("DROP TABLE IF EXISTS `%s`", table)); err != nil {
			gtest.Error(err)
		}
	}
}

func TestLogin(t *testing.T) {
	createTableFromSql()
	defer dropAllTableWithDb()

	// 登录成功
	gtest.C(t, func(t *gtest.T) {
		res, err := controller.Login{}.Login(context.Background(), &v1.LoginReq{
			Username: "admin",
			Password: "admin",
		})
		t.AssertNil(err)
		t.AssertNE(res.Token, "")
	})

	// 登录失败
	gtest.C(t, func(t *gtest.T) {
		res, err := controller.Login{}.Login(context.Background(), &v1.LoginReq{
			Username: "coding",
			Password: "coding123",
		})
		t.Assert(res, nil)
		t.AssertNE(err, nil)
	})
}

func TestRefresh(t *testing.T) {
	createTableFromSql()
	defer dropAllTableWithDb()

	res, err := controller.Login{}.Login(context.Background(), &v1.LoginReq{
		Username: "admin",
		Password: "admin",
	})
	if err != nil {
		t.Error(err)
	}
	token := res.Token
	// 正确的
	gtest.C(t, func(t *gtest.T) {
		ctx := context.Background()
		userClaim, err := logic.JwtHandler.Parse(ctx, token)
		t.Assert(err, nil)
		setCtx := func(key, value interface{}) {
			ctx = context.WithValue(ctx, key, value)
		}
		logic.CtxHandler.SetUserContext(userClaim.Username, setCtx)
		res, err := controller.Login{}.Refresh(ctx, &v1.RefreshReq{})
		t.AssertNil(err)
		t.AssertNE(res.Token, "")
	})

	// 错误的
	gtest.C(t, func(t *gtest.T) {
		ctx := context.Background()
		res, err := controller.Login{}.Refresh(ctx, &v1.RefreshReq{})
		t.AssertNE(err, nil)
		t.AssertNil(res)
	})

}
