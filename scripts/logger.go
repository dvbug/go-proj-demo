// Package core
// Copyright (C) Beijing Xiaomi Co., Ltd
// By Vito on 2022/1/18 11:31

package scripts

import "fmt"

type LogLevel int32

const (
	OFF LogLevel = iota
	ON
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
)

func (lv LogLevel) String() string {
	switch lv {
	case OFF:
		return "OFF"
	case ON:
		return "ON"
	case TRACE:
		return "T"
	case DEBUG:
		return "D"
	case INFO:
		return "I"
	case WARN:
		return "W"
	case ERROR:
		return "E"
	default:
		return "UNKNOWN"
	}
}

func printLn(str string, level LogLevel) {
	fmt.Println(fmt.Sprintf("LOG/%s %s", level.String(), str))
}

func Trace(str string) {
	printLn(str, TRACE)
}

func Debug(str string) {
	printLn(str, DEBUG)
}

func Info(str string) {
	printLn(str, INFO)
}

func Warn(str string) {
	printLn(str, WARN)
}

func Error(str string) {
	printLn(str, ERROR)
}
