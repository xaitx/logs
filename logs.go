package logs

import (
	"fmt"
	"io"
	"log"
	"os"
)

// 定义一个日志struct
type Logger struct {
	level    int         // 输出日志级别
	filePath string      // 输出日志路径
	logger   *log.Logger // 日志对象
	time     bool        // 是否开启日志前面的时间显示
}

// 定义日志级别常量
const (
	DEBUG = iota // DEBUG=0 灰色
	INFO         // INFO=1 绿色
	WARN         // WARN=2 黄色
	ERROR        // ERROR=3 红色
	FATAL        // FATAL=4 紫色
)

// NewLogger 创建一个新的日志记录器。
// 参数:
//   - level: 日志级别。
//   - filePath: 日志文件路径。如果为空，则日志输出到控制台。
//   - time: 是否在日志中包含时间。
//
// 返回值:
//   - *Logger: 新创建的日志记录器。
//   - error: 如果创建日志记录器时发生错误，则返回错误信息。
func NewLogger(level int, filePath string, time bool) (*Logger, error) {
	var flag int
	if time {
		flag = log.LstdFlags
	} else {
		flag = 0
	}

	logger := log.New(os.Stdout, "", flag)

	// 判断filePath是否为空，为空则输出到控制台
	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		logger.SetOutput(io.MultiWriter(file, os.Stdout))
	}

	return &Logger{
		level:    level,
		filePath: filePath,
		logger:   logger,
		time:     time,
	}, nil
}

// Debug 输出debug日志
func (l *Logger) Debug(v ...interface{}) {
	if l.level <= DEBUG {
		l.logger.SetPrefix("[DEBUG] ")
		l.logger.Println(Gray(fmt.Sprint(v...)))
	}
}

// Info 输出info日志
func (l *Logger) Info(v ...interface{}) {
	if l.level <= INFO {
		l.logger.SetPrefix("[INFO]  ")
		l.logger.Println(Green(fmt.Sprint(v...)))
	}
}

// Warn 输出warn日志
func (l *Logger) Warn(v ...interface{}) {
	if l.level <= WARN {
		l.logger.SetPrefix("[WARN]  ")
		l.logger.Println(Yellow(fmt.Sprint(v...)))
	}
}

// Error 输出error日志
func (l *Logger) Error(v ...interface{}) {
	if l.level <= ERROR {
		l.logger.SetPrefix("[ERROR] ")
		l.logger.Println(Red(fmt.Sprint(v...)))
	}
}

// Fatal 输出fatal日志
func (l *Logger) Fatal(v ...interface{}) {
	if l.level <= FATAL {
		l.logger.SetPrefix("[FATAL] ")
		l.logger.Fatalln(Purple(fmt.Sprint(v...)))
	}
}

// 输出不做任何处理的信息
func (l *Logger) Print(v ...interface{}) {
	l.logger.Println(v...)
}
