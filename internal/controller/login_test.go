package controller_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/stretchr/testify/assert"

	v1 "login-demo/api/v1"
	"login-demo/internal/controller"

	_ "login-demo/internal/packed"
)

func init() {
	// ctx := gctx.New()
	// User.InitAdmin(ctx)
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

func TestLogin(t *testing.T) {
	client := gclient.New()
	params := map[string]string{
		"username": "test_user",
		"password": "test_password",
	}
	// 发送 HTTP POST 请求
	resp, err := client.Post(context.Background(), "http://127.0.0.1:8000/api/user/login", params)
	assert.Empty(t, err)
	defer resp.Close()
	assert.Equal(t, resp.StatusCode, 200)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(resp.ReadAllString()), &data)
	// data := resp.ReadAllString()
	// .GetMap("data")
	assert.NotEmpty(t, data["token"].(string))
	// assert.Equal(t, data["role"].(string)), "test_role")
	// assert.Equal(t, data["role"].(string)), "test_role")
}

func TestLogin2(t *testing.T) {

	res, err := controller.Login{}.Login(context.Background(), &v1.LoginReq{
		Username: "admin",
		Password: "admin",
	})
	fmt.Println(res)
	fmt.Println(err)

}
