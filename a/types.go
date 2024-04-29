package a

import (
	"fmt"
	"strconv"
)

type Map = map[string]any
type M = map[string]any
type MS = map[string]string

type List = []any
type L = []any

type Void struct{}

var empty Void

type Function func(any) any
type Predicate func(any) bool
type Consumer func(any)
type Supplier func() any

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
