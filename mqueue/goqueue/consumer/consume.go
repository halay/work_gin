package consumer

import (
	"context"
	"github.com/zeromicro/go-zero/core/service"
	"work_gin/mqueue/goqueue/consumer/logic"
)

func InitConsumer() error {
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	mqs := logic.NewKqueue(context.Background())
	for _, mq := range mqs.Register() {
		serviceGroup.Add(mq)
	}
	serviceGroup.Start()
	return nil
}
