package response

const (
	UserExist             = 41100
	ErrAuth               = 41101
	ErrUsernameOrPassword = 41102
	ErrJwtToken           = 41103
	ErrAuthExpired        = 41104
	ErrAuthUnknown        = 41105
)

func authRes(msg map[int]string) map[int]string {
	msg[UserExist] = "用户已存在"
	msg[ErrAuth] = "未登录或非法访问"
	msg[ErrUsernameOrPassword] = "用户名或密码错误"
	msg[ErrJwtToken] = "创建token失败"
	msg[ErrAuthExpired] = "授权已过期"
	msg[ErrAuthUnknown] = "登陆未知错误"
	return msg
}
