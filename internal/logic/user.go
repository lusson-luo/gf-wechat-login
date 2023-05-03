package logic

import (
	"context"
	"errors"
	"fmt"
	"log"
	"login-demo/internal/dao"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"
	"strings"

	"crypto/sha256"

	"github.com/fatih/color"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

type LogicUser struct {
}

var (
	User = LogicUser{}
)

// Login 登录
func (lu *LogicUser) Login(ctx context.Context, username string, password string) (role string, token string, err error) {
	// 查询用户信息
	user, err := dao.User.Ctx(ctx).One("passport=? and password=?", username, fmt.Sprintf("%x", sha256.Sum256([]byte(password))))
	if err != nil {
		return "", "", err
	}
	if user == nil {
		err = errors.New("账户或密码错误")
		return "", "", err
	}
	// var user *entity.User
	// err = dao.User.Ctx(ctx).Where(do.User{
	// 	Passport: username,
	// 	Password: password,
	// }).Scan(&user)
	// if err != nil {
	// 	return
	// }
	// if user == nil {
	// 	err = errors.New("账户或密码错误")
	// 	return
	// }
	// 生成 jwt token
	token, err = MyJwt.GenerateToken(ctx, username)
	if err != nil {
		return "", "", err
	}
	// todo: 暂时没有 role
	return "", token, nil
}

// IsSignedIn 检查是否已经登录
func (lu *LogicUser) IsSignedIn(ctx context.Context, r *ghttp.Request) bool {
	token, exist := lu.getToken(r)
	if !exist {
		return false
	}
	valid := MyJwt.Valid(r.Context(), token)
	return valid
}

// Parse 解析 jwt token
func (lu *LogicUser) Parse(ctx context.Context, r *ghttp.Request) (bool, string) {
	token, exist := lu.getToken(r)
	if !exist {
		return false, ""
	}
	claims, ok := MyJwt.Parse(r.Context(), token)
	if !ok {
		return false, ""
	}
	return ok, claims.Username
}

// getToken 从 request 的 header 中获取 token
func (*LogicUser) getToken(r *ghttp.Request) (string, bool) {
	header := r.GetHeader("Authorization")
	headerList := strings.Split(header, " ")
	if len(headerList) != 2 {
		return "", false
	}
	t := headerList[0]
	token := headerList[1]
	if t != "Bearer" {
		return "", false
	}
	if token == "" {
		return "", false
	}
	return token, true
}

type AdminInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// InitAdmin 初始化管理员账户，从配置文件中
func (*LogicUser) InitAdmin(ctx context.Context) {
	admin := &AdminInfo{}
	err := g.Cfg().MustGet(ctx, "admin").Scan(admin)
	if err != nil {
		log.Fatal("读取admin配置失败")
		return
	}
	var user *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{
		Passport: admin.Username,
	}).Scan(&user)
	if err != nil {
		g.Log().Infof(ctx, color.RedString("err=%v, dao.User=%v"), err, user)
		panic("查询用户表失败，是否没有连接正确的数据库")
	}
	if user == nil {
		dao.User.Ctx(ctx).Insert(do.User{
			Passport: admin.Username,
			Password: fmt.Sprintf("%x", sha256.Sum256([]byte(admin.Password))),
			Nickname: admin.Username,
		})
	}
}

func init() {
	ctx := gctx.New()
	User.InitAdmin(ctx)
}
