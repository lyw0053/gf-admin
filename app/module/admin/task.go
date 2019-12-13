package admin

import "github.com/gogf/gf/net/ghttp"

type TaskController struct {
}

/**
模版渲染
*/
func (c *TaskController) Index(r *ghttp.Request) {
	r.Response.WriteTpl("task/index.html")
}

/**
任务数据列表
*/
func (c *TaskController) List(r *ghttp.Request) {
	r.Response.Write()
}
