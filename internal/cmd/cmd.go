package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	managerController "login-demo/internal/controller/manager"
	saasController "login-demo/internal/controller/saas"
	wxController "login-demo/internal/controller/wx"
	"login-demo/internal/logic"
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
				group.Group("/saas-api", func(group *ghttp.RouterGroup) {
					// 绑定 Login 结构体中的 Login 方法
					group.GET("", new(managerController.TenantController).SelectList)
					group.POST("", new(managerController.TenantController).Register)
					group.Middleware(middleware.TenantMiddleware)
					group.POST("", new(saasController.Login).Login)
					group.Group("", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.Auth)
						group.POST("", new(saasController.Login).Refresh)
						group.Bind(
							&saasController.UserController{},
							&saasController.StationController{},
							&saasController.PileController{},
							&saasController.ChargeOrderController{},
							&saasController.ChargePriceController{},
							&saasController.UserPayController{},
							&saasController.RoleController{},
							&saasController.PermissionController{},
							&saasController.AuditLogController{},
						)
					})
				})
				group.Group("/manager-api", func(group *ghttp.RouterGroup) {
					// 绑定 Login 结构体中的 Login 方法
					group.POST("", new(managerController.Login).Login)
					group.Group("", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.Auth)
						group.POST("", new(managerController.Login).Refresh)
						group.GET("", new(managerController.TenantController).List)
						group.POST("", new(managerController.TenantController).Add)
						group.POST("", new(managerController.TenantController).Update)
						group.POST("", new(managerController.TenantController).Del)
						group.Bind(
							&saasController.UserController{},
							&saasController.StationController{},
							&saasController.PileController{},
							&saasController.ChargeOrderController{},
							&saasController.ChargePriceController{},
							&saasController.UserPayController{},
							&saasController.RoleController{},
							&saasController.PermissionController{},
							&saasController.AuditLogController{},
						)
					})
				})
				// 微信路由
				group.Bind(
					&wxController.WxLoginController{},
				)
				group.GET("", new(wxController.WxChargeController).StationList)
				group.GET("", new(wxController.WxChargeController).PileList)
				group.Group("", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)
					group.POST("", new(wxController.WxChargeController).StartCharge)
					group.POST("", new(wxController.WxChargeController).MyChargeOrders)
					group.POST("", new(wxController.WxChargeController).AboutMe)
					group.POST("", new(wxController.WxChargeController).StopCharge)
					group.POST("", new(wxController.WxChargeController).PriceList)
				})
			})
			logic.User.InitAdmin(ctx)
			s.Run()
			return nil
		},
	}
)
