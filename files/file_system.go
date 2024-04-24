package files

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

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func IsDir(rawPath string) bool {
	stat, err := os.Stat(rawPath)
	if err != nil {
		return false
	}
	return stat.IsDir()
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

func FileList2(rawPath string) []*FileInfo2 {
	entries, err := os.ReadDir(rawPath)
	if err != nil {
		panic(err)
	}
	var res []*FileInfo2
	for _, entry := range entries {
		res = append(res, &FileInfo2{
			Dir:    entry.IsDir(),
			Parent: rawPath,
			Name:   entry.Name(),
			Path:   filepath.Join(rawPath, entry.Name()),
		})
	}
	return res
}

type FileInfo2 struct {
	Dir    bool   `json:"dir"`
	Parent string `json:"parent"`
	Name   string `json:"name"`
	Path   string `json:"path"`
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
func (receiver *FileInfo) Name() string {
	return receiver.raw.Name()
}
func (receiver *FileInfo) Dir() string {
	return receiver.baseDir
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
