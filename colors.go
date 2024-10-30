package logs

// 多样式的字体颜色输出

// Red 输出红色字体
func Red(str string) string {
	return "\x1b[31m" + str + "\x1b[0m"
}

// Green 输出绿色字体
func Green(str string) string {
	return "\x1b[32m" + str + "\x1b[0m"
}

// Yellow 输出黄色字体
func Yellow(str string) string {
	return "\x1b[33m" + str + "\x1b[0m"
}

// Blue 输出蓝色字体
func Blue(str string) string {
	return "\x1b[34m" + str + "\x1b[0m"
}

// Gray 输出灰色字体
func Gray(str string) string {
	return "\x1b[90m" + str + "\x1b[0m"
}

// Magenta 输出品红色字体
func Magenta(str string) string {
	return "\x1b[35m" + str + "\x1b[0m"
}

// Purple 输出紫色字体
func Purple(str string) string {
	return "\x1b[35m" + str + "\x1b[0m"
}

// Cyan 输出青色字体
func Cyan(str string) string {
	return "\x1b[36m" + str + "\x1b[0m"
}

// White 输出白色字体
func White(str string) string {
	return "\x1b[37m" + str + "\x1b[0m"
}

// Black 输出黑色字体
func Black(str string) string {
	return "\x1b[30m" + str + "\x1b[0m"
}

// Bold 输出加粗字体
func Bold(str string) string {
	return "\x1b[1m" + str + "\x1b[0m"
}
