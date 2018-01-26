package models

const (
	STATE_REALLOC int = 202 //重新分配
	STATE_CREATED int = 201 //初始创建
	STATE_STARTED int = 200 //成功状态
	STATE_STOPED  int = 0   //停止状态
	STATE_FAILED  int = -1  //失败状态
)

//GetStateString is exported
//return state string.
func GetStateString(state int) string {

	switch state {
	case STATE_REALLOC:
		return "STATE_REALLOC"
	case STATE_CREATED:
		return "STATE_CREATED"
	case STATE_STARTED:
		return "STATE_STARTED"
	case STATE_STOPED:
		return "STATE_STOPED"
	case STATE_FAILED:
		return "STATE_FAILED"
	}
	return ""
}
