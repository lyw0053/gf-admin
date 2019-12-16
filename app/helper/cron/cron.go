package cron

import "github.com/gogf/gf/os/gcron"

type Option struct {
	Pattern string
	Name    string
	Url     string
	Params  string
}

func CreateTask(op *Option) {
	gcron.Add(op.Pattern, func() {

	}, op.Name)
}
