package controller_wx_test

import (
	"context"
	v2 "login-demo/api/wx"
	"login-demo/internal/consts"
	controller "login-demo/internal/controller/wx"
	"login-demo/internal/model"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestAboutMe(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	// 正常
	gtest.C(t, func(t *gtest.T) {
		userCtx := context.WithValue(ctx, consts.UserContextKey, model.UserContext{
			Username: "admin",
		})
		res, err := controller.WxUserController{}.AboutMe(userCtx, &v2.WXMeInfoReq{})
		t.AssertNil(err)
		t.Assert(res.Username, "admin")
	})
	// 查询 guest 的订单
	gtest.C(t, func(t *gtest.T) {
		userCtx := context.WithValue(ctx, consts.UserContextKey, model.UserContext{
			Username: "guest",
		})
		res, err := controller.WxUserController{}.AboutMe(userCtx, &v2.WXMeInfoReq{})
		t.AssertNil(err)
		t.Assert(res.Username, "guest")
	})
	// 登录用户不存在
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.WxUserController{}.AboutMe(ctx, &v2.WXMeInfoReq{})
		t.AssertNE(err, nil)
	})
}
