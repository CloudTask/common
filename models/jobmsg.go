package models

import (
	"sync"
	"time"
)

/*
消息头类型定义
*/
const (
	MsgJobExecute = "JobExecute" //任务执行状态结果消息
	MsgJobSelect  = "JobSelect"  //任务选择下次执行消息
	//...更过消息预留扩展
)

/*
消息头定义
*/
type MsgHeader struct {
	MsgName string `json:"msgname"` //消息名称
	MsgId   string `json:"msgid"`   //消息编号
}

/*
任务执行状态结果消息体定义
Header.MsgName = MsgJobExecute
*/
type JobExecute struct {
	MsgHeader           //消息头
	JobId     string    `json:"jobid"`     //任务编号
	Location  string    `json:"location"`  //所在位置
	Key       string    `json:"key"`       //执行服务器
	IPAddr    string    `json:"ipaddr"`    //服务器地址
	State     int       `json:"state"`     //执行编码
	ExecErr   string    `json:"execerr"`   //执行错误信息
	ExecAt    time.Time `json:"execat"`    //本次执行时间
	NextAt    time.Time `json:"nextat"`    //下次执行时间
	Timestamp int64     `json:"timestamp"` //消息时间戳
}

type JobSelect struct {
	MsgHeader           //消息头
	JobId     string    `json:"jobid"`     //任务编号
	Location  string    `json:"location"`  //所在位置
	NextAt    time.Time `json:"nextat"`    //下次执行时间
	Timestamp int64     `json:"timestamp"` //消息时间戳
}

type MessageCache struct {
	sync.Mutex
	Data map[string]int64
}

func NewMessageCache() *MessageCache {

	return &MessageCache{
		Data: make(map[string]int64, 0),
	}
}

func (mc *MessageCache) ValidateMessage(message interface{}) bool {

	switch msgval := message.(type) {
	case *JobExecute:
		return mc.pressMessage(msgval.MsgName+"-"+msgval.JobId, msgval.Timestamp)
	case *JobSelect:
		return mc.pressMessage(msgval.MsgName+"-"+msgval.JobId, msgval.Timestamp)
	}
	return false
}

func (mc *MessageCache) pressMessage(key string, timestamp int64) bool {

	mc.Lock()
	defer mc.Unlock()
	if value, ret := mc.Data[key]; ret {
		if value > timestamp {
			return false
		}
	}
	mc.Data[key] = timestamp
	return true
}
