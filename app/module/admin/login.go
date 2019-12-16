package admin

import (
	"gf-admin/app/helper/request"
	"gf-admin/app/model"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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
		"username@required|phone#用户名不能为空|请填写正确的手机号",
		"password@required#密码不能为空",
	}

	if e := gvalid.CheckMap(params, rule); e != nil {
		request.ReturnJson(r, 1, e.FirstString(), nil)
		return
	}
	r.Session.Set("loginUserInfo", params)
	request.ReturnJson(r, 0, "登录成功", nil)
	return
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
