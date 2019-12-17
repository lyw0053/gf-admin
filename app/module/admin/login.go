package admin

import (
	"gf-admin/app/helper/request"
	"gf-admin/app/model"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gvalid"
)

type LoginController struct {
}

var (
	adminModel = (*model.Admin)(nil)
)

func (c *LoginController) Index(r *ghttp.Request) {
	r.Response.WriteTpl("login.html")
}

func (c *LoginController) Login(r *ghttp.Request) {
	params := g.Map{
		"username": r.GetForm("username", ""),
		"password": r.GetForm("password", ""),
		"code":     r.GetForm("code", 0),
	}
	rule := []string{
		"username@required#用户名不能为空",
		"password@required#密码不能为空",
	}

	if e := gvalid.CheckMap(params, rule); e != nil {
		request.ReturnJson(r, 1, e.FirstString(), nil)
		return
	}
	adminInfo, e := model.FindOne(adminModel, model.WithWhere(g.Map{"username": params["username"]}))
	if e != nil {
		request.ReturnJson(r, 1, e.Error(), nil)
		return
	}
	if pwd, _ := gmd5.Encrypt(params["password"]); pwd != adminInfo["password"] {
		request.ReturnJson(r, 1, "密码错误", nil)
		return
	}

	r.Session.Set("loginUserInfo", adminInfo)
	request.ReturnJson(r, 0, "登录成功", nil)

}

func (c *LoginController) LoginOut(r *ghttp.Request) {
	r.Session.Remove("loginUserInfo")
	request.ReturnJson(r, 0, "注销成功", nil)
}

/**
重置管理员账号
*/
func (c *LoginController) ResetAdmin(r *ghttp.Request) {
	admin, e := model.FindOne(adminModel, model.WithWhere(g.Map{"username": "admin"}))
	if e != nil {
		r.Response.Write(e)
		return
	}
	password, _ := gmd5.Encrypt("123456")
	if _, ok := admin["id"]; !ok {
		_, e := model.Insert(adminModel, model.WithData(g.Map{
			"username": "admin",
			"password": password,
			"nickname": "sinner",
			"add_time": gtime.Now(),
			"status":   1,
		}))
		if e != nil {
			r.Response.Write(e)
			return
		}
	} else {
		_, e := model.Update(adminModel, model.WithData(g.Map{"password": password}), model.WithWhere(g.Map{"id": admin["id"]}))
		if e != nil {
			r.Response.Write(e)
			return
		}
	}
	r.Response.Write("重置成功")
}
