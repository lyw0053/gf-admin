package request

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

/**
返回json
*/
func ReturnJson(r *ghttp.Request, code int, msg string, data interface{}) {
	rel := g.Map{}
	rel["code"] = code
	rel["msg"] = msg
	if data != nil {
		maps := data.(g.Map)
		if maps["count"] != nil {
			rel["data"] = maps["data"]
			rel["count"] = maps["count"]
		} else {
			rel["data"] = data
		}
	}

	r.Response.WriteJson(rel)
}

/**
返回json
*/
func ReturnXml(r *ghttp.Request, data interface{}) {
	r.Response.WriteXml(data)
}

/**
layout模板渲染
*/
func LayoutView(r *ghttp.Request, tpl string, params g.Map) {
	r.Response.WriteTpl("common/layout.html", g.Map{"contentTpl": tpl, "data": params})
}
