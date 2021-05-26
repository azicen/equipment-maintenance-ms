package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	//LevelError 错误
	LevelError = iota
	//LevelWarning 警告
	LevelWarning
	//LevelInformational 提示
	LevelInformational
	//LevelDebug 除错
	LevelDebug
)

var level int

//Println 打印
func Println(msg string) {
	log.Println(msg)
}

//Panic 极端错误
func Panic(format string, v ...interface{}) {
	if LevelError > level {
		return
	}
	msg := fmt.Sprintf("[FATAL] "+format, v...)
	Println(msg)
	panic(msg)
	//os.Exit(0)
}

//Error 错误
func Error(format string, v ...interface{}) {
	if LevelError > level {
		return
	}
	msg := fmt.Sprintf("[ERROR] "+format, v...)
	Println(msg)
}

//Warning 警告
func Warning(format string, v ...interface{}) {
	if LevelWarning > level {
		return
	}
	msg := fmt.Sprintf("[WARN] "+format, v...)
	Println(msg)
}

//Info 信息
func Info(format string, v ...interface{}) {
	if LevelInformational > level {
		return
	}
	msg := fmt.Sprintf("[INFO] "+format, v...)
	Println(msg)
}

//Debug 校验
func Debug(format string, v ...interface{}) {
	if LevelDebug > level {
		return
	}
	msg := fmt.Sprintf("[DEBUG] "+format, v...)
	Println(msg)
}

//BuildLogger 构建logger
func BuildLogger(l string) {
	level = LevelError
	switch l {
	case "error":
		level = LevelError
	case "warning":
		level = LevelWarning
	case "info":
		level = LevelInformational
	case "debug":
		level = LevelDebug
	}
}

func Init() {
	//读取配置文件，设置日志级别
	BuildLogger(os.Getenv("LOG_LEVEL"))
}
