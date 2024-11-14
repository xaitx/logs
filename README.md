# golang简单日志模块

## 简介
这是一个用Go语言编写的简单日志模块，旨在提供基本的日志记录功能。

## 例子

``` golang
package logs_test

import (
	"testing"

	"github.com/xaitx/logs"
)

func TestColorOutput(t *testing.T) {
	logs.PrintColor(logs.Red, "Red")
	logs.PrintColor(logs.Green, "Green")
	logs.PrintColor(logs.Yellow, "Yellow")
	logs.PrintColor(logs.Blue, "Blue")
	logs.PrintColor(logs.Gray, "Gray")
	logs.PrintColor(logs.Magenta, "Magenta")
	logs.PrintColor(logs.Purple, "Purple")
	logs.PrintColor(logs.Cyan, "Cyan")
	logs.PrintColor(logs.White, "White")
	logs.PrintColor(logs.Black, "Black")
	logs.PrintColor(logs.Bold, "Bold")
}

// 测试Logs输出
func TestLogs(t *testing.T) {
	logs.Debug("This is a debug message")
	logs.Info("This is an info message")
	logs.Warn("This is a warning message")
	logs.Error("This is an error message")
	logs.Print("This is a print message")
	logs.Fatal("This is a fatal message")
}
```
