package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func init() {
	//模版引擎配置
	v := g.View()
	v.AddPath("template")
	//log文件配置
	c := g.Config()
	logPath := c.GetString("setting.logpath")
	glog.SetPath(logPath)
	glog.SetStdoutPrint(true)
	//web server 配置
	s := g.Server()
	s.SetServerRoot("public")
	s.SetLogPath(logPath)
	s.SetNameToUriType(ghttp.URI_TYPE_ALLLOWER)
	s.SetPort(8199)
}
