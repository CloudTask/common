package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

type JobBase struct {
	JobId         string         `json:"jobid"`
	JobName       string         `json:"jobname"`
	FileName      string         `json:"filename"`
	FileCode      string         `json:"filecode"`
	Cmd           string         `json:"cmd"`
	Env           []string       `json:"env"`
	Timeout       int            `json:"timeout"`
	Version       int            `json:"version"`
	Schedule      []*Schedule    `json:"schedule"`
	NotifySetting *NotifySetting `json:"notifysetting"`
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

func ParseJobBase(data map[string]interface{}) (*JobBase, error) {

	value, ret := data["data"]
	if !ret {
		return nil, fmt.Errorf("jobbase prase %s", "data key invalid.")
	}

	if reflect.TypeOf(value).Kind() != reflect.Map {
		return nil, fmt.Errorf("jobbase prase %s", "data type invalid.")
	}

	value, ret = value.(map[string]interface{})["jobbase"]
	if !ret {
		return nil, fmt.Errorf("jobbase prase %s", "jobbase key invalid.")
	}

	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(value); err != nil {
		return nil, err
	}

	jobbase := &JobBase{}
	if err := json.NewDecoder(buf).Decode(jobbase); err != nil {
		return nil, err
	}

	if jobbase.Env == nil {
		jobbase.Env = []string{}
	}

	if jobbase.Schedule == nil {
		jobbase.Schedule = []*Schedule{}
	}
	return jobbase, nil
}

/*
func ParseJobBase(data map[string]interface{}) (*JobBase, error) {

	if v, ret := data["data"]; ret {
		value := v.(map[string]interface{})["jobbase"].(map[string]interface{})
		jobbase := &JobBase{
			JobId:         value["jobid"].(string),
			JobName:       value["jobname"].(string),
			Group:         value["group"].(string),
			FileName:      value["filename"].(string),
			FileCode:      value["filecode"].(string),
			Cmd:           value["cmd"].(string),
			Env:           parseEnv(value),
			Timeout:       (int)(value["timeout"].(float64)),
			Version:       (int)(value["version"].(float64)),
			Schedule:      parseSchedule(value),
			NotifySetting: parseNotifySetting(value),
		}
		return jobbase, nil
	}
	return nil, errors.New("read jobbase data error.")
}

func parseEnv(v map[string]interface{}) []string {

	env := []string{}
	if value, ret := v["env"]; ret {
		if value != nil && reflect.TypeOf(value).Kind() == reflect.Slice {
			env = value.([]string)
		}
	}
	return env
}

func parseSchedule(v map[string]interface{}) []*Schedule {

	schedules := []*Schedule{}
	if value, ret := v["schedule"]; ret {
		array := value.([]interface{})
		for _, it := range array {
			m := it.(map[string]interface{})
			schedule := &Schedule{
				Id:        m["id"].(string),
				Enabled:   (int)(m["enabled"].(float64)),
				TurnMode:  (int)(m["turnmode"].(float64)),
				Interval:  (int)(m["interval"].(float64)),
				StartDate: m["startdate"].(string),
				EndDate:   m["enddate"].(string),
				StartTime: m["starttime"].(string),
				EndTime:   m["endtime"].(string),
				SelectAt:  m["selectat"].(string),
			}
			if monthlyof, ok := m["monthlyof"]; ok {
				pvalue := monthlyof.(map[string]interface{})
				schedule.MonthlyOf.Day = (int)(pvalue["day"].(float64))
				schedule.MonthlyOf.Week = pvalue["week"].(string)
			}
			schedules = append(schedules, schedule)
		}
	}
	return schedules
}

func parseNotifySetting(v map[string]interface{}) *NotifySetting {

	if value, ret := v["notifysetting"]; ret {
		successed := value.(map[string]interface{})["succeed"].(map[string]interface{})
		failed := value.(map[string]interface{})["failed"].(map[string]interface{})
		return &NotifySetting{
			Succeed: Notify{
				Enabled: successed["enabled"].(bool),
				Subject: successed["subject"].(string),
				To:      successed["to"].(string),
				Content: successed["content"].(string),
			},
			Failed: Notify{
				Enabled: failed["enabled"].(bool),
				Subject: failed["subject"].(string),
				To:      failed["to"].(string),
				Content: failed["content"].(string),
			},
		}
	}
	return nil
}
*/
