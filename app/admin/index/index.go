package index

import "github.com/gogf/gf/net/ghttp"

type Controller struct {
}

func (c *Controller) Index(r *ghttp.Request) {
	r.Response.WriteTpl("index.html")
}
