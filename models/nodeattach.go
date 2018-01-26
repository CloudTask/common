package models

import (
	"bytes"
	"encoding/json"
	"sync"
)

var attach_pool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 1<<10))
	},
}

/*
节点附加数据定义
*/
type AttachData struct {
	JobMaxCount int `json:"jobmaxcount"` //每个节点的最大任务数配置
}

func AttachEncode(attach *AttachData) []byte {

	buffer := attach_pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer attach_pool.Put(buffer)
	if err := json.NewEncoder(buffer).Encode(attach); err != nil {
		return []byte{}
	}
	return buffer.Bytes()
}

func AttachDecode(data []byte) *AttachData {

	attach := &AttachData{}
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(attach); err != nil {
		return &AttachData{}
	}
	return attach
}
