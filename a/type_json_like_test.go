package a

import "testing"

func Test_Arr(t *testing.T) {
	arr := ArrNew()
	arr.A(1, 9, "油荒", "测试", true)

	Echo(arr.Json())
	F("-%v-\n", arr.Int(9))
	F("-%v-\n", arr.Int(1))
	F("-%v-\n", arr.Int(2))
	F("-%v-\n", arr.Str(2))
	F("-%v-\n", arr.Str(1))
	Echo("-------------------")
	arr.Set(0, M{"a": "bbbbbbbbbbbbbbb"})
	arr.Set(1, L{"a", "b", "c", "d", "e"})
	getArr := arr.GetArr(1)
	Echo(getArr.Json())
	getArr2 := arr.GetArr(2)
	Echo(getArr2.Json())
	obj := arr.GetObj(0)
	Echo(obj.Json())
	obj2 := arr.GetObj(1)
	Echo(obj2.Json())
}

func Test_Obj55(t *testing.T) {
	obj := ObjNew()
	obj.S("k", 9).
		S("v", Now().String()).
		S("obj", M{"name": "天氧！"})
	Echo(obj.Json())
	F("-%v-\n", obj.Int("v"))
	F("-%v-\n", obj.Str("k"))
	F("-%v-\n", obj.Str("v"))
	Echo("=============================")
	obj.S("map", M{"action": "竞争者"})
	obj.S("list", L{"actor", "lamba", false, 998})
	o := obj.GetObj("map")
	o2 := obj.GetObj("list")
	a := obj.GetArr("map")
	a2 := obj.GetArr("list")
	Echo(o.Json())
	Echo(o2.Json())
	Echo(a.Json())
	Echo(a2.Json())
}
