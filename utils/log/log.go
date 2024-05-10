package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"work_gin/utils"
)

var Log *logrus.Logger

func InitLog() {
	filePath := utils.LogPath + "log"
	//linkName := "latest_log.log"

	scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err:", err)
	}
	Log = logrus.New()

	Log.Out = scr
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Log.SetLevel(logrus.DebugLevel)
}
