package f

import (
	"github.com/changlie/go-common/a"
	"testing"
)

func Test_append(t *testing.T) {
	fpath := "/home/gx/ws/xxx"
	file := OpenAppendFile(fpath)
	defer file.Close()

	file.WriteLnChar()
	file.WriteStringLn("瑶池阿母绮窗开")
	file.WriteLn([]byte("黄竹歌声动地哀"))
	file.WriteStringLn("八骏日行三万里")
	file.WriteLn([]byte("穆王何事不重来"))

	for i, s := range ReadLines(fpath) {
		a.Fln("[%v] -> %v", i, s)
	}
}
