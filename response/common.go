package response

var codeMsgMap map[int]string

// 错误码
const (
	CodeSuccess = 0
	// 4 开头的是前端操作问题
	ErrCodeParameter = 41001
	ErrAuthExoired   = 41003
	ErrAuthUnknown   = 41004
	// 5 开头是后端问题
	ErrSQL = 52001
)

func Init() {
	codeMsgMap = make(map[int]string, 1024)
	codeMsgMap = baseRes(codeMsgMap)
	codeMsgMap = authRes(codeMsgMap)
}

func baseRes(msg map[int]string) map[int]string {
	msg[CodeSuccess] = "success"
	msg[ErrCodeParameter] = "参数错误"
	msg[ErrSQL] = "sql错误"
	msg[ErrAuthExoired] = "授权已过期"
	msg[ErrAuthUnknown] = "登陆未知错误"
	return msg
}

func getMessage(code int) (message string) {
	message, ok := codeMsgMap[code]
	if !ok {
		message = "未知错误"
	}
	return message
}
