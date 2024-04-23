package a

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func Uuid() string {
	return strings.ReplaceAll(UuidV4(), "-", "")
}

func UuidV4() string {
	// V4 基于随机数
	u4 := uuid.New()
	return u4.String()
}

func ProgramDir() string {
	if executable, err := os.Executable(); err == nil {
		return filepath.Dir(executable)
	} else {
		panic(err)
		return ""
	}
}

func Exec(command string) string {
	return doCmd(command)
}

func doCmd(command string) string {
	var cmder *exec.Cmd
	if runtime.GOOS == "windows" {
		cmder = exec.Command("cmd", "/C", command)
	} else {
		cmder = exec.Command("bash", "-c", command)
	}

	output, err := cmder.CombinedOutput()
	if err != nil {
		fmt.Println(err, cmdResult(output))
		return ""
	}
	return cmdResult(output)
}

func cmdResult(bs []byte) string {
	bs = bytes.TrimSpace(bs)
	if runtime.GOOS != "windows" {
		return string(bs)
	}

	resBytes, err := simplifiedchinese.GBK.NewDecoder().Bytes(bs)
	if err != nil {
		return string(bs)
	} else {
		return string(resBytes)
	}
}
