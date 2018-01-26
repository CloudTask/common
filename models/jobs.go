package models

import (
	"time"
)

/*
 任务轮询模式信息
*/
const (
	TURNMODE_SECONDS = 1 //秒轮询
	TURNMODE_MINUTES = 2 //分钟轮询
	TURNMODE_HOURLY  = 3 //小时轮询
	TURNMODE_DAILY   = 4 //天轮询
	TURNMODE_WEEKLY  = 5 //周轮询
	TURNMODE_MONTHLY = 6 //月轮询
)

/*
 月轮询选项
*/
type MonthlyOf struct {
	/* Day 选定月份第n天
	   > 0 表示每月第n天
	   < 0 表示每月倒数第n天
	   == 0 表示按Week方式处理
	*/
	Day int `json:"day"`
	/*
	   Week 选定月份第n个星期, 星期由0~6编码表示
	   "0:1" 表示最后一个星期一
	   "1:2" 表示第一个星期二
	   "2:4" 表示第二个星期四
	*/
	Week string `json:"week"`
}

/*
 计划信息
*/
type Schedule struct {
	Id        string    `json:"id"`        //计划编号
	Enabled   int       `json:"enabled"`   //计划开关(0:关 1:开)
	TurnMode  int       `json:"turnmode"`  //轮询模式
	Interval  int       `json:"interval"`  //轮询间隔
	StartDate string    `json:"startdate"` //开始日期
	EndDate   string    `json:"enddate"`   //结束日期
	StartTime string    `json:"starttime"` //开始时间
	EndTime   string    `json:"endtime"`   //结束时间
	SelectAt  string    `json:"selectat"`  //选择条件(条件已","符号分隔)
	MonthlyOf MonthlyOf `json:"monthlyof"` //月轮询选项
}

/*
 任务信息
*/
type Job struct {
	JobId         string         `json:"jobid"`         //任务编号
	Name          string         `json:"name"`          //任务名称
	Location      string         `json:"location"`      //隶属位置
	Servers       []string       `json:"servers"`       //分配服务器信息，若为空则由调度器以hash分配，一但设置了Servers，Job仅限于Servers范围进行分配.
	GroupId       string         `json:"groupid"`       //隶属分组
	FileName      string         `json:"filename"`      //文件名称
	Cmd           string         `json:"cmd"`           //执行命令
	Env           []string       `json:"env"`           //环境变量
	Timeout       int            `json:"timeout"`       //执行超时(0:永远等待)
	Enabled       int            `json:"enabled"`       //任务开关(0:关 1:开)
	Schedule      []*Schedule    `json:"schedule"`      //任务计划
	NotifySetting *NotifySetting `json:"notifysetting"` //任务通知配置
	Stat          int            `json:"stat"`          //执行编码
	ExecErr       string         `json:"execerr"`       //执行错误信息
	ExecAt        time.Time      `json:"execat"`        //本次执行时间
	NextAt        time.Time      `json:"nextat"`        //下次执行时间
}

/*
 任务基础简要信息
 用于任务调度等信息交换
*/
type SimpleJob struct {
	JobId    string   `json:"jobid"`    //任务编号
	Name     string   `json:"name"`     //任务名称
	Location string   `json:"location"` //隶属位置
	GroupId  string   `json:"groupid"`  //隶属分组
	Servers  []string `json:"servers"`  //分配服务器信息
	Enabled  int      `json:"enabled"`  //任务开关(0:关 1:开)
	Stat     int      `json:"stat"`     //执行编码
}
