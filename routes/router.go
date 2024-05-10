package routes

import (
	"github.com/gin-gonic/gin"
	"work_gin/api/yto"
	"work_gin/middleware"
	"work_gin/utils"
	"work_gin/utils/log"
)

func InitRouter() error {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(middleware.RateLimit())
	/*
		前端展示页面接口
	*/
	router := r.Group("api/yto")
	{
		//测试
		//router.GET("test/redis/string", yto.RedisString)
		//router.GET("test/redis/hash", yto.RedisHash)
		router.POST("audio/download", yto.DownloadMP3Handler)
		//router.GET("rate/test", middleware.RateLimit(), yto.RateTest)
		router.GET("asynq/test", yto.AsynqTest)
	}
	log.Log.Info("service start success....")
	err := r.Run(utils.HttpPort)
	if err != nil {
		log.Log.Infof("service start err,%+v \n", err)
		return err
	}
	return nil
}
