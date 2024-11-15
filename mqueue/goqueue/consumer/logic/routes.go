package logic

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"work_gin/utils"
)

type Kqueue struct {
	ctx context.Context
}

func NewKqueue(ctx context.Context) *Kqueue {
	return &Kqueue{
		ctx: ctx,
	}
}
func (k *Kqueue) Register() []service.Service {
	return []service.Service{
		kq.MustNewQueue(utils.KqConf, NewSendMsg()),
	}
}
