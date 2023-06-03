package a

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"log"
)

func Json(raw any) string {
	bytes, err := json.Marshal(raw)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func JsonOf(raw string) *Kson {
	newJson, err := simplejson.NewJson([]byte(raw))
	if err != nil {
		log.Println(err)
		return nil
	}
	return &Kson{raw: newJson}
}

func FromJson(raw string, res any) {
	if err := json.Unmarshal([]byte(raw), res); err != nil {
		log.Fatal(err)
	}
}

type Kson struct {
	raw *simplejson.Json
}

func (receiver *Kson) Get(key string) *Kson {
	tmp := receiver.raw.Get(key)
	return &Kson{raw: tmp}
}
func (receiver *Kson) Set(key string, val any) {
	receiver.raw.Set(key, val)
}
func (receiver *Kson) GetIndex(index int) *Kson {
	tmp := receiver.raw.GetIndex(index)
	return &Kson{raw: tmp}
}
func (receiver *Kson) Encode() string {
	bytes, err := receiver.raw.Encode()
	if err != nil {
		return ""
	}
	return string(bytes)
}
func (receiver *Kson) EncodePretty() string {
	bytes, err := receiver.raw.EncodePretty()
	if err != nil {
		return ""
	}
	return string(bytes)
}
func (receiver *Kson) Model(res any) {
	bytes, err := receiver.raw.Encode()
	if err != nil {
		return
	}
	if err := json.Unmarshal(bytes, res); err != nil {
		log.Fatal(err)
	}
}

func (receiver *Kson) Int(key string) int {
	i, err := receiver.raw.Get(key).Int()
	if err != nil {
		log.Println(err)
		return 0
	}
	return i
}
func (receiver *Kson) Int64(key string) int64 {
	i, err := receiver.raw.Get(key).Int64()
	if err != nil {
		return 0
	}
	return i
}
func (receiver *Kson) Float64(key string) float64 {
	f, err := receiver.raw.Get(key).Float64()
	if err != nil {
		return 0
	}
	return f
}
func (receiver *Kson) String(key string) string {
	s, err := receiver.raw.Get(key).String()
	if err != nil {
		return ""
	}
	return s
}
func (receiver *Kson) Bytes(key string) []byte {
	bytes, err := receiver.raw.Get(key).Bytes()
	if err != nil {
		return nil
	}
	return bytes
}
func (receiver *Kson) Bool(key string) bool {
	flag, err := receiver.raw.Get(key).Bool()
	if err != nil {
		return false
	}
	return flag
}
func (receiver *Kson) ArrLen() int {
	array, err := receiver.raw.Array()
	if err != nil {
		return -1
	}
	return len(array)
}
func (receiver *Kson) Array(key string) []any {
	array, err := receiver.raw.Get(key).Array()
	if err != nil {
		return nil
	}
	return array
}
func (receiver *Kson) Map(key string) map[string]any {
	m, err := receiver.raw.Get(key).Map()
	if err != nil {
		return nil
	}
	return m
}
func (receiver *Kson) Val(key string) any {
	return receiver.raw.Get(key).Interface()
}

func (receiver *Kson) IntByIndex(index int) int {
	i, err := receiver.raw.GetIndex(index).Int()
	if err != nil {
		log.Println(err)
		return 0
	}
	return i
}
func (receiver *Kson) Int64ByIndex(index int) int64 {
	i, err := receiver.raw.GetIndex(index).Int64()
	if err != nil {
		return 0
	}
	return i
}
func (receiver *Kson) Float64ByIndex(index int) float64 {
	f, err := receiver.raw.GetIndex(index).Float64()
	if err != nil {
		return 0
	}
	return f
}
func (receiver *Kson) StringByIndex(index int) string {
	s, err := receiver.raw.GetIndex(index).String()
	if err != nil {
		return ""
	}
	return s
}
func (receiver *Kson) BytesByIndex(index int) []byte {
	bytes, err := receiver.raw.GetIndex(index).Bytes()
	if err != nil {
		return nil
	}
	return bytes
}
func (receiver *Kson) BoolByIndex(index int) bool {
	flag, err := receiver.raw.GetIndex(index).Bool()
	if err != nil {
		return false
	}
	return flag
}
func (receiver *Kson) ArrayByIndex(index int) []any {
	array, err := receiver.raw.GetIndex(index).Array()
	if err != nil {
		return nil
	}
	return array
}
func (receiver *Kson) MapByIndex(index int) map[string]any {
	m, err := receiver.raw.GetIndex(index).Map()
	if err != nil {
		return nil
	}
	return m
}
func (receiver *Kson) ValByIndex(index int) any {
	return receiver.raw.GetIndex(index).Interface()
}
