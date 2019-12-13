package admin

import (
	"gf-admin/app/helper/request"
	"gf-admin/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type TaskController struct {
}

var (
	taskModel = (*model.Task)(nil)
)

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
	where := g.Map{}
	page := r.GetInt("page", 1)
	limit := r.GetInt("limit", 10)
	offset := (page - 1) * limit
	id := r.GetInt("id", 0)
	if id > 0 {
		where["id"] = id
	}
	name := r.GetString("name", "")
	if name != "" {
		where["name like ?"] = "%" + name + "%"
	}
	count, _ := model.Count(taskModel, model.WithWhere(where))
	if count == 0 {
		request.ReturnJson(r, 0, "succ", g.Map{"data": nil, "count": 0})
	}
	lists, e := model.FindAll(taskModel, model.WithWhere(where), model.WithLimit([]int{offset, limit}))
	if e != nil {
		request.ReturnJson(r, 1, e.Error(), nil)
		return
	}
	request.ReturnJson(r, 0, "succ", g.Map{"data": lists, "count": count})

}
