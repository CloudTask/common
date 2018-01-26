package models

/*
消息头类型定义
*/
const (
	MsgSystemEvent = "SystemEvent" //系统事件消息
)

/*
事件类型定义
*/
const (
	CreateJobEvent      = "create_job"      //创建任务事件
	ChangeJobEvent      = "change_job"      //修改任务事件
	RemoveJobEvent      = "remove_job"      //删除任务事件
	ChangeJobsFileEvent = "change_jobsfile" //任务文件批量修改事件
	CreateGroupEvent    = "create_group"    //创建分组事件
	ChangeGroupEvent    = "change_group"    //改变分组事件
	RemoveGroupEvent    = "remove_group"    //删除分组事件
)

/*
系统事件定义
*/
type SystemEvent struct {
	MsgHeader          //消息头
	Event     string   `json:"event"`     //事件名称
	Runtime   string   `json:"runtime"`   //runtime名称
	JobIds    []string `json:"jobids"`    //事件相关所有jobs
	GroupIds  []string `json:"groupids"`  //事件相关所有groups
	Timestamp int64    `json:"timestamp"` //消息时间戳
}
