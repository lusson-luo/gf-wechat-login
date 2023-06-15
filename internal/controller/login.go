package controller

import (
	"context"
	v1 "login-demo/api/v1"
	"login-demo/internal/logic"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// var Login CLogin = CLogin{}

type Login struct {
}

// 登录
func (c Login) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	role, tokenString, err := logic.User.Login(ctx, req.Username, req.Password)
	if err != nil {
		return
	}
	res = &v1.LoginRes{
		Token: tokenString,
		Role:  role,
	}
	return
}

func (C Login) Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error) {
	user, isExist := logic.CtxHandler.GetUserContext(ctx)
	if !isExist {
		err = gerror.NewCode(gcode.New(1, "设置 ctx 缓存失败，系统异常", "设置 ctx 缓存失败，系统异常"))
		return
	}
	newToken, err := logic.JwtHandler.GenerateToken(ctx, user.Username)
	res = &v1.RefreshRes{
		Token: newToken,
	}
	return
}
