package middleware

import (
	"github.com/gogf/gf/net/ghttp"
)

func LoginCheck(r *ghttp.Request) {
	userInfo := r.Session.Get("loginUserInfo")
	if userInfo == nil {
		r.Response.RedirectTo("/admin/login")
	}
	r.Middleware.Next()
}
