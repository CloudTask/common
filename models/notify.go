package models

/*
 通知定义
*/
type Notify struct {
	Enabled bool   `json:"enabled"` //通知开关
	Subject string `json:"subject"` //通知标题
	To      string `json:"to"`      //接收者
	Content string `json:"content"` //通知内容
}

/*
 通知设置定义
*/
type NotifySetting struct {
	Succeed Notify `json:"succeed"` //成功通知
	Failed  Notify `json:"failed"`  //失败通知
}
