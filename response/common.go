package response

var codeMsgMap map[int]string

// 错误码
const (
	CodeSuccess = 0
	// 4 开头的是前端操作问题
	ErrCodeParameter = 41001
	ErrAuth          = 41002
	ErrAuthExoired   = 41003
	ErrAuthUnknown   = 41004
	// 5 开头是后端问题
	ErrSQL = 52001
)

func Init() {
	codeMsgMap = make(map[int]string, 1024)
	codeMsgMap[CodeSuccess] = "success"
	codeMsgMap[ErrCodeParameter] = "参数错误"
	codeMsgMap[ErrSQL] = "sql错误"
	codeMsgMap[ErrAuth] = "未登录或非法访问"
	codeMsgMap[ErrAuthExoired] = "授权已过期"
	codeMsgMap[ErrAuthUnknown] = "登陆未知错误"
}

func getMessage(code int) (message string) {
	message, ok := codeMsgMap[code]
	if !ok {
		message = "未知错误"
	}
	return message
}
