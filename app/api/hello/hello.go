package hello

import (
	"gf-admin/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Hello World
func Handler(r *ghttp.Request) {
	results, e := model.FindAll((*model.User)(nil), model.WithWhere(g.Map{"id": 1}))
	if e != nil {
		r.Response.Writeln(e)
		return
	}
	r.Response.Writeln(results)
}
