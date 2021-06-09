package response

const (
	UserExist = 41100
	ErrAuth   = 41002
)

func authRes(msg map[int]string) map[int]string {
	msg[UserExist] = "用户已存在"
	msg[ErrAuth] = "未登录或非法访问"
	return msg
}
