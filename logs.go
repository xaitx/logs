package logs

import (
	"fmt"
	"io"
	"log"
	"os"
)

// 定义一个日志struct
type Logger struct {
	Level    int         // 输出日志级别
	FilePath string      // 输出日志路径
	logger   *log.Logger // 日志对象
}

// 定义日志级别常量
const (
	DEBUG = iota // DEBUG=0 灰色
	INFO         // INFO=1 绿色
	WARN         // WARN=2 黄色
	ERROR        // ERROR=3 红色
	FATAL        // FATAL=4 紫色
)

// NewLogger 创建一个日志对象
func NewLogger(level int, filePath string) (*Logger, error) {

	logger := log.New(os.Stdout, "", log.LstdFlags)

	// 判断filePath是否为空，为空则输出到控制台
	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		logger.SetOutput(io.MultiWriter(file, os.Stdout))
	}

	return &Logger{
		Level:    level,
		FilePath: filePath,
		logger:   logger,
	}, nil
}

// Debug 输出debug日志
func (l *Logger) Debug(v ...interface{}) {
	if l.Level <= DEBUG {
		l.logger.SetPrefix("[DEBUG] ")
		l.logger.Println(Gray(fmt.Sprint(v...)))
	}
}

// Info 输出info日志
func (l *Logger) Info(v ...interface{}) {
	if l.Level <= INFO {
		l.logger.SetPrefix("[INFO]  ")
		l.logger.Println(Green(fmt.Sprint(v...)))
	}
}

// Warn 输出warn日志
func (l *Logger) Warn(v ...interface{}) {
	if l.Level <= WARN {
		l.logger.SetPrefix("[WARN]  ")
		l.logger.Println(Yellow(fmt.Sprint(v...)))
	}
}

// Error 输出error日志
func (l *Logger) Error(v ...interface{}) {
	if l.Level <= ERROR {
		l.logger.SetPrefix("[ERROR] ")
		l.logger.Println(Red(fmt.Sprint(v...)))
	}
}

// Fatal 输出fatal日志
func (l *Logger) Fatal(v ...interface{}) {
	if l.Level <= FATAL {
		l.logger.SetPrefix("[FATAL] ")
		l.logger.Fatalln(Purple(fmt.Sprint(v...)))
	}
}
