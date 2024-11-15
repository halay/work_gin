package yto

import (
	"github.com/gin-gonic/gin"
	"time"
	"work_gin/model"
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
func HTest(c *gin.Context) {
	//fmt.Println("rate test" + time.Now().String())
	//time.Sleep(time.Second * 10)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "halay test" + time.Now().String(),
	})
}
