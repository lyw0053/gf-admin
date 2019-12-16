package router

import (
	"gf-admin/app/middleware"
	"gf-admin/app/module/admin"
	"gf-admin/app/module/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 统一路由注册.
func init() {
	s := g.Server()
	s.Group("/admin", func(g *ghttp.RouterGroup) {
		g.ALL("/login", new(admin.LoginController))
		g.Middleware(middleware.LoginCheck)
		g.ALL("/", new(admin.IndexController))
		g.ALL("/task", new(admin.TaskController))
	})
	s.Group("/api", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.Auth)
		g.ALL("/cron", new(api.CronController))
	})
}
