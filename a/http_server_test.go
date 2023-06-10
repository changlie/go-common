package a

import (
	"fmt"
	"testing"
)

type Users2 struct {
	Name string `json:"name,omitempty" form:"name"`
	Pwd  string `json:"pwd,omitempty" form:"pwd"`
}

func Fox(c Ctx) {
	var u Users2
	c.JsonArgs(&u)
	c.Ok(M{
		"action": "天极",
		"age":    10086,
		"args":   u,
		"at":     Now().String(),
	})
}

func Test_simple(t *testing.T) {
	srv := HttpServerNew()
	srv.Post("/fox", Fox)
	srv.Get("/tt", Tget)
	srv.Start()
}

func Tget(c Ctx) {
	var u Users2
	c.QueryArgs(&u)
	headers := c.HeaderMap()
	for k, v := range headers {
		fmt.Printf("header(%v -> %v)\n", k, v)
	}
	c.Status(404).JsonBody(M{
		"coed": 999999,
		"msg":  "这是个无法到达的页面！",
		"args": u,
		"hs":   headers,
	})
}
