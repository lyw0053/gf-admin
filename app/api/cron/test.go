package cron

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
)

type Controller struct {
}

func (this *Controller) Index(r *ghttp.Request) {
	cronName := "myTestCron"
	status := r.GetPostString("status", "start")

	switch status {
	case "start":
		if gcron.Search(cronName) != nil {
			gcron.Start(cronName)
			r.Response.Write("任务启动!")
			return
		} else {
			gcron.Add("* * * * * *", func() {
				glog.Println("秒级定时任务：" + gtime.Datetime())
			}, cronName)
			r.Response.Write("任务首次创建并启动!")
			return
		}

	case "stop":
		if gcron.Search(cronName) != nil {
			gcron.Stop(cronName)
			r.Response.Write("任务关闭！")
			return
		}
		r.Response.Write("任务未开启!")
		return
	}
}

func (this *Controller) Test(r *ghttp.Request) {
	r.Response.Write("test")
}
