package files

import (
	"fmt"
	"github.com/changlie/go-common/a"
	"testing"
)

func Test_file_list2(t *testing.T) {
	list := FileList2("D:/")
	fmt.Println(a.JsonOf(list).EncodePretty())
}
