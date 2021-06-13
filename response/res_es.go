package response

const (
	EsNotEnable   = 52000
	EsSearchError = 52001
)

func esRes(msg map[int]string) map[int]string {
	msg[EsNotEnable] = "配置没有打开es"
	msg[EsSearchError] = "查询日志失败"
	return msg
}
