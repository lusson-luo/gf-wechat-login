package logic

import (
	"context"
	"fmt"
	"login-demo/internal/model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type LogicCtxHandler struct {
}

var (
	CtxHandler = LogicCtxHandler{}
)

const (
	UserContextKey = "userCtx"
)

type UserContext struct {
	Username string
}

// 设置 username 到 context 中
func (LogicCtxHandler) SetUserContext(username string, setCtx func(interface{}, interface{})) {
	setCtx(UserContextKey, UserContext{
		Username: username,
	})
}

// 从 context 取出 user 信息
func (LogicCtxHandler) GetUserContext(ctx context.Context) (*UserContext, bool) {
	v, ok := ctx.Value(UserContextKey).(UserContext)
	return &v, ok
}

func (l LogicCtxHandler) GetCurrentUser(ctx context.Context) (user model.UserMore, err error) {
	userCtx, _ := l.GetUserContext(ctx)
	users, count, err := User.UserList(ctx, userCtx.Username, model.PageReq{})
	if err != nil {
		return
	}
	if count != 1 {
		err = gerror.NewCode(gcode.New(404, fmt.Sprintf("%s 用户不存在，请重新注册", userCtx.Username), "系统异常"))
		return
	}
	user = users[0]
	return
}
