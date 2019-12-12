package user

import (
	"gf-admin/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
}

var (
	userModel = (*model.User)(nil)
)

func (c *Controller) List(r *ghttp.Request) {
	users, e := model.FindAll(userModel, model.WithLimit([]int{10}))
	if e != nil {
		r.Response.Writeln(e)
		return
	}
	r.Response.Writeln(users)
}

func (c *Controller) One(r *ghttp.Request) {
	uid := r.GetInt("uid", 0)
	if uid == 0 {
		r.Response.Write("无效的uid")
		return
	}
	user, e := model.FindOne(userModel, model.WithWhere(g.Map{"id": uid}))
	if e != nil {
		r.Response.Writeln(e)
		return
	}
	r.Response.Writeln(user)
}
