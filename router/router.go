package router

import (
	"gf-admin/app/admin/index"
	"gf-admin/app/api/cron"
	"gf-admin/app/api/hello"
	"gf-admin/app/api/user"
	"gf-admin/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 统一路由注册.
func init() {
	g.Server().BindHandler("/", hello.Handler)

	s := g.Server()
	s.Group("/admin", func(g *ghttp.RouterGroup) {
		g.ALL("/", new(index.Controller))
	})
	s.Group("/api", func(g *ghttp.RouterGroup) {
		g.ALL("/user", new(user.Controller))
		g.Middleware(middleware.Auth)
		g.ALL("/cron", new(cron.Controller))
	})
}
