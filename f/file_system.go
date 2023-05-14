package f

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func Rm(rawPath string) {
	err := os.RemoveAll(rawPath)
	if err != nil {
		panic(err)
	}
}

func Mkdir(rawPath string) {
	err := os.MkdirAll(rawPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func FileList(rawPath string) []*FileInfo {
	entries, err := os.ReadDir(rawPath)
	if err != nil {
		panic(err)
	}
	var res []*FileInfo
	for _, entry := range entries {
		res = append(res, NewFileInfo(rawPath, entry))
	}
	return res
}

type FileInfo struct {
	baseDir string
	raw     os.DirEntry
}

func NewFileInfo(base string, rawFileInfo os.DirEntry) *FileInfo {
	return &FileInfo{baseDir: base, raw: rawFileInfo}
}

func (receiver *FileInfo) Path() string {
	return filepath.Join(receiver.baseDir, receiver.raw.Name())
}

func (receiver *FileInfo) IsDir() bool {
	return receiver.raw.IsDir()
}

func (receiver *FileInfo) Type() fs.FileMode {
	return receiver.raw.Type()
}

func (receiver *FileInfo) Info() fs.FileInfo {
	info, err := receiver.raw.Info()
	if err != nil {
		panic(fmt.Sprintf("%v: %v", receiver.Path(), err))
	}
	return info
}

func (receiver *FileInfo) Files() []*FileInfo {
	if receiver.raw.IsDir() {
		return FileList(receiver.Path())
	}
	return nil
}
