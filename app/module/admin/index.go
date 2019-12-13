package admin

import "github.com/gogf/gf/net/ghttp"

type IndexController struct {
}

/**
index
*/
func (c *IndexController) Index(r *ghttp.Request) {
	r.Response.WriteTpl("index.html")
}

/**
后台首页
*/
func (c *IndexController) Home(r *ghttp.Request) {
	r.Response.WriteTpl("home.html")
}
