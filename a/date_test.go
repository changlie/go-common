package a

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println(Now().DateStr())
}

func Test2(t *testing.T) {
	fmt.Println(Now())
	fmt.Println(Now().AddHour(-5).AddMinute(-15).AddSecond(-7))
	fmt.Println(Now().Format("y/M/d"))
	fmt.Println(Now().AddYear(2).AddDay(9).Format("y年M月d日 h时m分s秒"))
	fmt.Println(Now().DateStr())
}
