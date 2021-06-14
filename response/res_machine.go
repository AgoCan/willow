package response

const (
	MachinePasswordIsNull   = 53000
	MachinePrivateKeyIsNull = 53001
	MachineNameExist        = 53002
	MachineHostExist        = 53003
	MachineTypeError        = 53004
)

func machineRes(msg map[int]string) map[int]string {
	msg[MachinePasswordIsNull] = "选择密码方式，而密码为空"
	msg[MachinePrivateKeyIsNull] = "选择密钥方式，而密钥为空"
	msg[MachineNameExist] = "节点名称已经存在"
	msg[MachineHostExist] = "主机已经存在"
	msg[MachineTypeError] = "类型错误"
	return msg
}