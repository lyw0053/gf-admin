package model

import "github.com/gogf/gf/os/gtime"

type Task struct {
	Id      int64      `orm:id`
	Name    string     `orm:name`
	Type    int        `orm:type`
	Content string     `orm:content`
	AddTime gtime.Time `orm:add_time`
	Status  int        `orm:status`
	CronId  string     `orm:cron_id`
}

const (
	//get
	TASK_TYPE_GET = 1
	//post
	TASK_TYPE_POST = 2
)

func (this *Task) GetTableName() string {
	return tablePrefix + "task"
}
