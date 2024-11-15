package logic

import (
	"encoding/json"
	"fmt"
	"work_gin/mqueue/goqueue/common"
	"work_gin/utils/log"
)

type SendMsg struct {
}

func NewSendMsg() *SendMsg {
	return &SendMsg{}
}
func (s *SendMsg) Consume(key, value string) error {
	log.Log.Info("SendMsg Consume key:", key, "value:", value)
	var m common.SendMsg
	if err := json.Unmarshal([]byte(value), &m); err != nil {
		log.Log.Error("SendMsg Consume json.Unmarshal err:", err)
		return err
	}
	fmt.Printf("send the msg %s to user %d success...", m.Msg, m.UserId)
	log.Log.Infof("send the msg %s to user %d success...", m.Msg, m.UserId)
	return nil
}
