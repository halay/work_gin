package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"work_gin/mqueue/asynq/common"
)

type SendEmail struct {
}

func NewSendEmail() *SendEmail {
	return &SendEmail{}
}
func (s *SendEmail) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p common.EmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	fmt.Println("this send email processTask success,and the userId is :", p.UserId)
	return nil
}
