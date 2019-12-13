package request

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func ReturnJson(r *ghttp.Request, code int, msg string, data interface{}) {
	rel := g.Map{}
	rel["code"] = code
	rel["msg"] = msg
	maps := data.(g.Map)
	if maps["count"] != nil {
		rel["data"] = maps["data"]
		rel["count"] = maps["count"]
	} else {
		rel["data"] = data
	}
	r.Response.Write(rel)
}
