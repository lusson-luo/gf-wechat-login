package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"login-demo/internal/controller"
	"login-demo/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					middleware.ParseJwtToCtx,
				)
				// 绑定 Login 结构体中的 Login 方法
				group.POST("/login", new(controller.Login).Login)
				group.Bind(
				// controller.Login,
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)
					group.POST("/refresh", new(controller.Login).Refresh)
					group.Bind(
					// controller.Login;
					)
				})
				// group.POST("/login", controller.Login.Login, middleware.ParseJwtToCtx)
				// group.POST("/refresh", controller.Login.Refresh, middleware.ParseJwtToCtx, middleware.Auth)
			})
			s.Run()
			return nil
		},
	}
)
