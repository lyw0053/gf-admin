package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func LoginCheck(r *ghttp.Request) {
	token := r.Get("token")
	if token != "123456" {
		r.Response.Write(g.Map{
			"code": 0,
			"msg":  "认证失败",
			"data": nil,
		})
		return
	}
	r.Response.Writeln("auth")
	r.Middleware.Next()
}
