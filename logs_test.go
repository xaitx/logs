package logs_test

import (
	"testing"

	"github.com/xaitx/logs"
)

func TestColorOutput(t *testing.T) {
	// 通过颜色输出日志
	logger, err := logs.NewLogger(logs.DEBUG, "", false)
	if err != nil {
		t.Error(err)
	}
	// 输出测试颜色
	logger.Print(logs.Gray("Gray"))
	logger.Print(logs.Green("Green"))
	logger.Print(logs.Yellow("Yellow"))
	logger.Print(logs.Red("Red"))
	logger.Print(logs.Purple("Purple"))
	logger.Print(logs.Cyan("Cyan"))
	logger.Print(logs.White("White"))
	logger.Print(logs.Black("Black"))
	logger.Print(logs.Blue("Blue"))
	logger.Print(logs.Magenta("Magenta"))
	logger.Print(logs.Cyan("Cyan"))

}

// 测试Logs输出
func TestLogs(t *testing.T) {
	logger, err := logs.NewLogger(logs.DEBUG, "", false)
	if err != nil {
		t.Error(err)
	}
	logger.Debug("this is debug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
	// logger.Fatal("this is fatal")
}
