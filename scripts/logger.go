// Package core
// Copyright (C) Vito
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

func printLn(v interface{}, level LogLevel) {
	fmt.Println(fmt.Sprintf("LOG/%s %v", level.String(), v))
}

func Trace(v interface{}) {
	printLn(v, TRACE)
}

func Debug(v interface{}) {
	printLn(v, DEBUG)
}

func Info(v interface{}) {
	printLn(v, INFO)
}

func Warn(v interface{}) {
	printLn(v, WARN)
}

func Error(v interface{}) {
	printLn(v, ERROR)
}
