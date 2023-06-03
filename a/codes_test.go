package a

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_hex(t *testing.T) {
	str := "Hello from ADMFactory.com"
	hx := hex.EncodeToString([]byte(str))
	fmt.Println("String to Hex Golang example")
	fmt.Println()
	fmt.Println(str + " ==> " + hx)
	bytes, err := hex.DecodeString(hx)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))

	raw := "中华人民共和国"
	encode := HexEncode(raw)
	fmt.Println("encode", encode)
	fmt.Println("decode", string(HexRaw(encode)))

}

func Test_base64(t *testing.T) {

}
