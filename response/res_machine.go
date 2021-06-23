package response

const (
	MachinePasswordIsNull   = 53000
	MachinePrivateKeyIsNull = 53001
	MachineNameExist        = 53002
	MachineHostExist        = 53003
	MachineTypeError        = 53004
	MachineNameNotExist     = 53005
	MachineGroupNotExist    = 53006
	MachineGroupExist       = 53007
	MachineGroupIDIsNull    = 53008
)

func machineRes(msg map[int]string) map[int]string {
	msg[MachinePasswordIsNull] = "选择密码方式，而密码为空"
	msg[MachinePrivateKeyIsNull] = "选择密钥方式，而密钥为空"
	msg[MachineNameExist] = "节点名称已经存在"
	msg[MachineHostExist] = "主机已经存在"
	msg[MachineTypeError] = "类型错误"
	msg[MachineNameNotExist] = "节点不存在"
	msg[MachineGroupNotExist] = "机器组不存在"
	msg[MachineGroupExist] = "机器组已存在"
	msg[MachineGroupIDIsNull] = "机器和机器组为空"
	return msg
}
