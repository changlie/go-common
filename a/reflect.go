package a

import (
	"fmt"
	"reflect"
	"strings"
)

type FunctionExecutor struct {
	name string         // 函数名称
	ins  []reflect.Type //入参类型
	outs []reflect.Type // 出参类型
	obj  reflect.Value  // 函数对象
}

func (f *FunctionExecutor) Run(args []reflect.Value) any {
	resList := f.obj.Call(args)
	if len(resList) < 1 {
		return nil
	}
	res := resList[0]
	return res.Interface()
}

func (f *FunctionExecutor) InNum() int {
	return len(f.ins)
}

func (f *FunctionExecutor) InLastIndex() int {
	return len(f.ins) - 1
}

func (f *FunctionExecutor) OutNum() int {
	return len(f.outs)
}

// 通过反射收集函数信息
func CollectFunctionInfo(objDoublePtr interface{}) (res map[string]*FunctionExecutor) {
	res = make(map[string]*FunctionExecutor)
	v1 := reflect.ValueOf(objDoublePtr).Elem()
	k1 := v1.Type()
	for i := 0; i < v1.NumMethod(); i++ {
		funcExe := &FunctionExecutor{}

		methodName := k1.Method(i).Name
		methodObject := v1.Method(i)

		methodType := methodObject.Type()
		// in params
		incount := methodType.NumIn()
		for ii := 0; ii < incount; ii++ {
			argType := methodType.In(ii)
			funcExe.ins = append(funcExe.ins, argType)
		}

		// out params
		outcount := methodType.NumOut()
		for ii := 0; ii < outcount; ii++ {
			argType := methodType.Out(ii)
			funcExe.outs = append(funcExe.outs, argType)
		}

		funcExe.obj = methodObject
		funcExe.name = formatName(methodName)
		res[funcExe.name] = funcExe
	}
	return res
}

// 将函数名第一个字母转小写
func formatName(methodName string) string {
	return fmt.Sprintf("%v%v", strings.ToLower(methodName[:1]), methodName[1:])
}
