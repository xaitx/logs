package logs

// 多样式的字体颜色输出

// Red 输出红色字体
func red(str string) string {
	return "\x1b[31m" + str + "\x1b[0m"
}

// Green 输出绿色字体
func green(str string) string {
	return "\x1b[32m" + str + "\x1b[0m"
}

// yellow 输出黄色字体
func yellow(str string) string {
	return "\x1b[33m" + str + "\x1b[0m"
}

// blue 输出蓝色字体
func blue(str string) string {
	return "\x1b[34m" + str + "\x1b[0m"
}

// gray 输出灰色字体
func gray(str string) string {
	return "\x1b[90m" + str + "\x1b[0m"
}

// magenta 输出品红色字体
func magenta(str string) string {
	return "\x1b[35m" + str + "\x1b[0m"
}

// purple 输出紫色字体
func purple(str string) string {
	return "\x1b[35m" + str + "\x1b[0m"
}

// cyan 输出青色字体
func cyan(str string) string {
	return "\x1b[36m" + str + "\x1b[0m"
}

// white 输出白色字体
func white(str string) string {
	return "\x1b[37m" + str + "\x1b[0m"
}

// black 输出黑色字体
func black(str string) string {
	return "\x1b[30m" + str + "\x1b[0m"
}

// bold 输出加粗字体
func bold(str string) string {
	return "\x1b[1m" + str + "\x1b[0m"
}

var (
	Red     = red     // 红色
	Green   = green   // 绿色
	Yellow  = yellow  // 黄色
	Blue    = blue    // 蓝色
	Gray    = gray    // 灰色
	Magenta = magenta // 品红色
	Purple  = purple  // 紫色
	Cyan    = cyan    // 青色
	White   = white   // 白色
	Black   = black   // 黑色
	Bold    = bold    // 加粗
)
