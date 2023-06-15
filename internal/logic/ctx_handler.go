package logic

import (
	"context"
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
