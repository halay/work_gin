package mqueue

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-queue/kq"
	"work_gin/mqueue/goqueue/common"
	"work_gin/utils"
)

func Push(c *gin.Context) {
	producer := kq.NewPusher(utils.KqConf.Brokers, utils.KqConf.Topic)
	var m = common.SendMsg{
		UserId: 1,
		Msg:    "hello halay",
	}
	body, err := json.Marshal(m)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "json marshal error",
		})
		return
	}
	err = producer.Push(string(body))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "push error" + err.Error(),
		})
		return
	}
}
