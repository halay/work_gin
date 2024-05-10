package yto

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"time"
	"work_gin/model"
	"work_gin/mqueue/asynq/common"
	"work_gin/utils"
	"work_gin/utils/log"
)

func RedisString(c *gin.Context) {
	model.RedisString()
}
func RedisHash(c *gin.Context) {
	model.RedisHash()
}
func RateTest(c *gin.Context) {
	//fmt.Println("rate test" + time.Now().String())
	//time.Sleep(time.Second * 10)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "rate test" + time.Now().String(),
	})
}
func AsynqTest(c *gin.Context) {
	asynqClient := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     utils.RedisArr,
		Password: utils.RedisPassword,
		DB:       1,
	})
	jobPayload, err := json.Marshal(common.EmailPayload{
		UserId: 1,
	})
	if err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "asynq err" + err.Error(),
		})
	}
	taskInfo, err := asynqClient.Enqueue(asynq.NewTask(common.JobSendEmail, jobPayload))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "asynq err" + err.Error(),
		})
	}
	log.Log.Info("asynq success,task info1 log ", taskInfo)

	taskInfo2, err := asynqClient.Enqueue(asynq.NewTask(common.JobSendEmail, jobPayload), asynq.ProcessIn(1*time.Minute))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "asynq err" + err.Error(),
		})
	}
	log.Log.Info("asynq success,task info2 log ", taskInfo2)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  fmt.Sprintf("asynq success, task info1: %v,and task info2 :%v", taskInfo, taskInfo2),
	})
}
