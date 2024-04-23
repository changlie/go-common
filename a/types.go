package a

import (
	"fmt"
	"strconv"
)

type Map = map[string]any
type M = map[string]any

type List = []any
type L = []any

type Void struct{}

var empty Void

type Set map[any]Void
type Consumer func(any)

func StrToInt(raw string) int {
	res, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return res
}
func StrToFloat(raw string) float64 {
	res, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		panic(err)
	}
	return res
}
func StrToFloat32(raw string) float32 {
	res, err := strconv.ParseFloat(raw, 32)
	if err != nil {
		panic(err)
	}
	return float32(res)
}
func StrToBool(raw string) bool {
	res, err := strconv.ParseBool(raw)
	if err != nil {
		panic(err)
	}
	return res
}
func IntStr(raw int) string {
	return fmt.Sprintf("%v", raw)
}
func FloatStr(raw float64, precision ...int) string {
	format := "%v"
	if len(precision) > 0 {
		format = fmt.Sprintf("%%.%vf", precision[0])
	}
	return fmt.Sprintf(format, raw)
}
func Float32Str(raw float32, precision ...int) string {
	format := "%v"
	if len(precision) > 0 {
		format = fmt.Sprintf("%%.%vf", precision[0])
	}
	return fmt.Sprintf(format, raw)
}
func BoolStr(raw bool) string {
	return fmt.Sprintf("%v", raw)
}

func NewSet(items ...any) Set {
	var res Set = make(map[any]Void)
	for _, item := range items {
		res.Add(item)
	}
	return res
}

func (s *Set) Add(item any) {
	(*s)[item] = empty
}
func (s *Set) Contains(item any) bool {
	_, ok := (*s)[item]
	return ok
}

func (s *Set) Del(item any) {
	delete(*s, item)
}

func (s *Set) Each(acceptor Consumer) {
	for k := range *s {
		acceptor(k)
	}
}

func (s *Set) ToArr() []any {
	var res []any
	for k := range *s {
		res = append(res, k)
	}
	return res
}
