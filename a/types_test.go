package a

import (
	"testing"
)

type Users struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Score        []int    `json:"score"`
	RoleCodeList []string `json:"roleCodeList"`
	UserAlias    string   `json:"userAlias"`
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

	ustr := `{"id":9, "name":"tom", "age":17, "score":[59,71,62]}`
	var u Users
	FromJson(ustr, &u)
	Echo(u.Name, Json(u))
}
