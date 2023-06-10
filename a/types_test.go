package a

import (
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
