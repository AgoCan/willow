package response

import "willow/model"

// Ping 测试序列化器
type Ping struct {
	ID  uint   `json:"id"`
	Msg string `json:"msg"`
}

//BuildPing 测试序列化器
func BuildPing(ping model.Ping) Ping {
	return Ping{
		ID:  ping.ID,
		Msg: ping.Msg,
	}
}
