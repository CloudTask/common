package models

import (
	"bytes"
	"encoding/json"
)

type JobBase struct {
	JobId    string      `json:"jobid"`
	JobName  string      `json:"jobname"`
	FileName string      `json:"filename"`
	FileCode string      `json:"filecode"`
	Cmd      string      `json:"cmd"`
	Env      []string    `json:"env"`
	Timeout  int         `json:"timeout"`
	Version  int         `json:"version"`
	Schedule []*Schedule `json:"schedule"`
}

func JobBaseEnCode(jobbase *JobBase) ([]byte, error) {

	buf := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buf).Encode(jobbase); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func JobBaseDeCode(buf []byte) (*JobBase, error) {

	jobbase := &JobBase{}
	r := bytes.NewReader(buf)
	if err := json.NewDecoder(r).Decode(jobbase); err != nil {
		return nil, err
	}
	return jobbase, nil
}
