package a

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func Test_reflect2(t *testing.T) {
	obj := &ReflectTest{}
	set := CollectFunctionInfo(&obj)
	for name, executor := range set {
		fmt.Println(name)

		if name == "now" {
			fmt.Println(executor.Run(nil))
		} else if name == "testStr" {
			params := []reflect.Value{
				reflect.ValueOf("hello"),
			}
			fmt.Println(executor.Run(params))
		} else {
			fmt.Println(executor.Run(nil))
		}
		fmt.Println("----------------")
	}
}

type ReflectTest struct {
}

func (r *ReflectTest) TestStr(arg string) any {
	return strings.ToUpper(arg)
}

func (r *ReflectTest) Now() any {
	return time.Now()
}

func (r *ReflectTest) Fow() any {
	return "你好，萨罗！"
}
