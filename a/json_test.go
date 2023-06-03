package a

import (
	"github.com/bitly/go-simplejson"
	"testing"
)

func Test_jsonLib1(t *testing.T) {
	js, _ := simplejson.NewJson([]byte(`{
        "test": {
            "string_array": ["asdf", "ghjk", "zxcv"],
            "string_array_null": ["abc", null, "efg"],
            "array": [1, "2", 3],
            "arraywithsubs": [{"subkeyone": 1},
            {"subkeytwo": 2, "subkeythree": 3}],
            "int": 10,
            "float": 5.150,
            "string": "simplejson",
            "bool": true,
            "sub_obj": {"a": 1}
        }
    }`))
	i, err := js.Get("test").Get("sub_obj").Get("a").Int()
	if err != nil {
		return
	}
	println(i)
	s, err := js.Get("test").Get("string_array").GetIndex(2).String()
	if err != nil {
		return
	}
	println(s)
}

func Test_customJsonLib1(t *testing.T) {
	raw := `{
        "test": {
            "string_array": ["asdf", "ghjk", "zxcv"],
            "string_array_null": ["abc", null, "efg"],
            "array": [1, "2", 3],
            "arraywithsubs": [{"subkeyone": 1},
            {"subkeytwo": 2, "subkeythree": 3}],
            "int": 10,
            "float": 5.150,
            "string": "simplejson",
            "bool": true,
            "sub_obj": {"a": 1},
			"stu":{"id":998,"name":"tome", "age":88, "sex":"猫"}
        }
    }`
	json := JsonOf(raw)
	arr := json.Get("test").Get("string_array")
	for i := 0; i < arr.ArrLen(); i++ {
		println(i, "->", arr.StringByIndex(i))
	}
	println(json.Get("test").Get("stu").EncodePretty())
}

type IdPair struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Test_customJsonLib2(t *testing.T) {
	raw := `
{"name":"天气伙",
"type":"free",
"stus":[
{"id":9,"name":"范"},
{"id":11,"name":"进益"}
]
}
`
	json := JsonOf(raw)
	var list []IdPair
	json.Get("stus").Model(&list)
	Echo(len(list), list[1].Name)
	list[1].Name = "无敌"
	list = append(list, IdPair{Id: 223, Name: "泰崔"})
	json.Set("stus", list)
	println(json.EncodePretty())
}
