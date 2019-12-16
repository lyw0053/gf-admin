package admin

import (
	"gf-admin/app/helper/request"
	"gf-admin/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gvalid"
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
	request.LayoutView(r, "task/index.html", nil)
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

func (c *TaskController) Add(r *ghttp.Request) {
	request.LayoutView(r, "task/add.html", nil)
}

func (c *TaskController) DoEdit(r *ghttp.Request) {
	formMap := r.GetFormMap()
	doedit, ok := formMap["doedit"]
	if !ok {
		request.ReturnJson(r, 1, "非法访问", nil)
		return
	}
	params := g.Map{
		"name":    r.GetForm("name", ""),
		"url":     r.GetForm("url", ""),
		"type":    r.GetForm("type", model.TASK_TYPE_GET),
		"content": r.GetForm("content", ""),
		"status":  r.GetFormInt("is_running", 0),
	}
	rule := []string{
		"name@required|max-length:150#任务名称不能为空|最大长度不能超过150字符",
		"url@required|max-length:200#url不能为空|最大长度不能超过200字符",
		"type@required",
		"status@boolean#is_running:无效参数",
	}

	if e := gvalid.CheckMap(params, rule); e != nil {
		request.ReturnJson(r, 1, e.FirstString(), nil)
		return
	}

	switch doedit.(string) {
	case "add":
		params["add_time"] = gtime.Now()
		params["status"] = 0
		taskId, e := model.Insert(taskModel, model.WithData(params))
		if e != nil {
			request.ReturnJson(r, 1, "添加失败", e)
			return
		}
		request.ReturnJson(r, 0, "添加成功", g.Map{"taskid": taskId})
	case "edit":
		request.ReturnJson(r, 0, "编辑成功", nil)
	default:
		request.ReturnJson(r, 1, "无效请求", nil)

	}

}
