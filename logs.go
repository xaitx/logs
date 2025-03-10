package logs

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

var (
	level  = DEBUG                                 // 输出日志级别
	logger = log.New(os.Stdout, "", log.LstdFlags) // 日志输出对象
	prefix = true                                  // 是否输出日志的前缀
)

// SetLevel 设置日志输出级别 DEBUG, INFO, WARN, ERROR, FATAL
func SetLevel(l int) {
	level = l
}

// SetOutput 设置日志输出对象,默认为os.Stdout
func SetOutput(w *os.File) {
	logger.SetOutput(w)
}

// SetFlags 设置日志输出格式
//
// 0 代表不输出时间，默认为log.LstdFlags eg: 2024/11/14 14:25:54
func SetFlags(flag int) {
	logger.SetFlags(flag)
}

// SetLogger 设置日志输出对象
//
// eg: logs.SetLogger(os.Stdout) 默认输出到控制台，也可以传入文件对象
func SetLogger(w io.Writer) {
	logger.SetOutput(w)
}

// 设置是否输出日志的前缀，比如INFO的前缀为[INFO], 默认为true
func SetPrefix(b bool) {
	prefix = b
}

// Debug 输出debug日志
func Debug(v ...interface{}) {
	if level > DEBUG {
		return
	}

	if prefix {
		logger.SetPrefix("[DEBUG] ")
	} else {
		logger.SetPrefix("")
	}

	logger.Println(Gray(fmt.Sprint(v...)))
}

// Info 输出info日志
func Info(v ...interface{}) {
	if level > INFO {
		return
	}

	if prefix {
		logger.SetPrefix("[INFO]  ")
	} else {
		logger.SetPrefix("")
	}

	logger.Println(Green(fmt.Sprint(v...)))
}

// Warn 输出warn日志
func Warn(v ...interface{}) {
	if level > WARN {
		return
	}

	if prefix {
		logger.SetPrefix("[WARN]  ")
	} else {
		logger.SetPrefix("")
	}

	logger.Println(Yellow(fmt.Sprint(v...)))
}

// Error 输出error日志
func Error(v ...interface{}) {
	if level > ERROR {
		return
	}

	if prefix {
		logger.SetPrefix("[ERROR] ")
	} else {
		logger.SetPrefix("")
	}

	logger.Println(Red(fmt.Sprint(v...)))
}

// Fatal 输出fatal日志
func Fatal(v ...interface{}) {
	if level > FATAL {
		return
	}

	if prefix {
		logger.SetPrefix("[FATAL] ")
	} else {
		logger.SetPrefix("")
	}

	logger.Fatalln(Purple(fmt.Sprint(v...)))
}

// 输出不做任何处理的信息
func Print(v ...interface{}) {
	logger.SetPrefix("")
	logger.Print(v...)
}

// 输出带颜色的信息
func PrintColor(x func(string) string, v ...interface{}) {
	logger.SetPrefix("")
	logger.Print(x(fmt.Sprint(v...)))
}
