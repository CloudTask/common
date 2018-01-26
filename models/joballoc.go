package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"sync"
)

/*
分配表单条Job信息
*/
type JobData struct {
	JobId   string `json:"jobid"`   //任务编号
	Key     string `json:"key"`     //任务名称
	Version int    `json:"version"` //任务版本
}

/*
任务分配表信息定义
*/
type AllocMapper map[string]*JobsAlloc

type JobsAlloc struct {
	Version int        `json:"version"` //分配表版本
	Jobs    []*JobData `json:"jobs"`    //任务列表
}

/*
编码分配表
*/
func JobsAllocEnCode(pool *sync.Pool, alloc *JobsAlloc) ([]byte, error) {

	buffer := pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer pool.Put(buffer)
	if err := json.NewEncoder(buffer).Encode(alloc); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

/*
解码分配表
*/
func JobsAllocDeCode(b []byte, alloc *JobsAlloc) error {

	if len(b) <= 0 {
		return errors.New("decode alloc buffer invalid.")
	}

	if err := json.NewDecoder(bytes.NewReader(b)).Decode(alloc); err != nil {
		return err
	}
	return nil
}

/*
Job分配信息
*/
type JobDataInfo struct {
	JobData
	JobName string `json:"jobname"`
	IpAddr  string `json:"ipaddr"`
}

/*
Job分配表详细数据
*/
type JobsAllocData struct {
	Location string         `json:"location"`
	Version  int            `json:"version"`
	Data     []*JobDataInfo `json:"data"`
}
