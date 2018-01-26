package models

import (
	"time"
)

/*
任务日志信息定义
*/
type JobLog struct {
	JobId     string    `json:"jobid"`     //任务编号
	MsgId     string    `json:"msgid"`     //消息编号
	Event     string    `json:"event"`     //事件描述
	Group     string    `json:"group"`     //隶属分组
	Stat      int       `json:"stat"`      //状态编码
	Location  string    `json:"location"`  //所在位置
	Command   string    `json:"command"`   //执行命令
	WorkDir   string    `json:"workdir"`   //执行工作目录
	IpAddr    string    `json:"ipaddr"`    //服务器地址
	StdOut    string    `json:"stdout"`    //标准输出
	ErrOut    string    `json:"errout"`    //错误输出
	ExecErr   string    `json:"execerr"`   //执行错误信息
	ExecAt    time.Time `json:"execat"`    //本次执行时间
	ExecTimes float64   `json:"exectimes"` //执行耗时(毫秒)
	CreateAt  int64     `json:"createat"`  //日志时间
}
