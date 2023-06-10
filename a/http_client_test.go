package a

import (
	"fmt"
	"testing"
)

func TestHttpGet(t *testing.T) {
	resp := HttpGet("https://api.nextrt.com/V1/Dutang")
	fmt.Println(string(resp))
}
