package a

import (
	"fmt"
	"reflect"
	"testing"
)

type Users struct {
	Id           int      `json:"id,omitempty" form:"id"`
	Name         string   `json:"name,omitempty" form:"name"`
	Age          int      `json:"age,omitempty" form:"age"`
	Score        []int    `json:"score,omitempty" form:"score"`
	RoleCodeList []string `json:"roleCodeList,omitempty" form:"roleCodeList"`
	UserAlias    string   `json:"userAlias,omitempty" form:"userAlias"`
}

func Test_MapList(t *testing.T) {
	obj := M{
		"name":   "changlie",
		"age":    99,
		"scores": L{1, 2, 3},
		"attr": M{
			"a": 1,
			"b": 2,
			"c": 3,
		},
	}
	println(Json(obj))

	blob := `["small","regular","large","unrecognized","small","normal","small","large"]`
	var arr []string
	FromJson(blob, &arr)
	Echo(len(arr), arr[3], Json(arr))

	ustr := `{"id":7,"name":"tom", "score":[59,71,62], "roleCodeList":["a","1"]}`
	var u Users
	FromJson(ustr, &u)
	Echo(u.Name, Json(u))
}

func TestNumberToStr(t *testing.T) {
	Echo(FloatStr(9.687777777777777777777777))
	Echo(FloatStr(9.687777777777777777777777, 2))
	Echo(FloatStr(9.687777777777777777777777, 3))
}

func TestStrToNumber(t *testing.T) {
	r1 := StrToInt("10087")
	Echo(r1, reflect.TypeOf(r1))
	r2 := StrToInt("123456789")
	Echo(r2, reflect.TypeOf(r2))
}

func Test_Set(t *testing.T) {
	set := NewSet("a", "C", "D", false, 998)
	set.Each(func(a any) {
		fmt.Println("item -> ", a)
	})
	Echo(set.Contains("A"))
	Echo(set.Contains("a"))
	Echo(set.Contains(998))
}
