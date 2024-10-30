package logs_test

import (
	"log"
	"testing"

	"github.com/xaitx/logs"
)

func TestColorOutput(t *testing.T) {
	// 通过颜色输出日志
	log.Println(logs.Red("this is red"))
	log.Println(logs.Green("this is green"))
	log.Println(logs.Yellow("this is yellow"))
	log.Println(logs.Blue("this is blue"))
	log.Println(logs.Gray("this is gray"))
	log.Println(logs.Magenta("this is magenta"))
	log.Println(logs.Purple("this is purple"))
	log.Println(logs.Cyan("this is cyan"))
	log.Println(logs.White("this is white"))
	log.Println(logs.Black("this is black"))
	log.Println(logs.Bold("this is bold"))
}

// 测试Logs输出
func TestLogs(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	logger, err := logs.NewLogger(logs.DEBUG, "E://tmp/logs.log")
	if err != nil {
		t.Error(err)
	}
	logger.Debug("this is debug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
	logger.Fatal("this is fatal")
}
