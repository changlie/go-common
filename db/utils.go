package db

import "fmt"

const (
	CommonDatetimeFormat = "2006-01-02 15:04:05"
	CommonDateFormat     = "2006-01-02"
)

func intToRunes(raw int) []rune {
	s := fmt.Sprint(raw)
	var res []rune
	for _, ch := range s {
		res = append(res, ch)
	}
	return res
}

func assert(flag bool, msg ...interface{}) {
	if flag {
		runtimeExcption(msg)
	}
}

// 报错并退出程序(带格式化)
func errorf(format string, args ...interface{}) {
	var msg []interface{}
	for _, item := range args {
		if err, ok := item.(error); ok && err != nil {
			msg = append(msg, err.Error())
			continue
		}
		msg = append(msg, item)
	}
	panic(fmt.Sprintf(format, msg...))
}

// 报错并退出程序(不带格式化)
func runtimeExcption(raw ...interface{}) {
	var msg []interface{}
	for _, item := range raw {
		if err, ok := item.(error); ok && err != nil {
			msg = append(msg, err.Error())
			continue
		}
		msg = append(msg, item)
	}
	panic(fmt.Sprint(msg...))
}
