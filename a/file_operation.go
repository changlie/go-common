package a

import (
	"bufio"
	"os"
)

func ReadBytes(rawPath string) []byte {
	fileBytes, err := os.ReadFile(rawPath)
	if err != nil {
		panic(err)
	}
	return fileBytes
}

func ReadString(rawPath string) string {
	fileBytes := ReadBytes(rawPath)
	return string(fileBytes)
}

func ReadLines(rawPath string) []string {
	file, err := os.Open(rawPath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func WriteBytes(rawPath string, data []byte) {
	err := os.WriteFile(rawPath, data, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
func WriteString(rawPath string, data string) {
	err := os.WriteFile(rawPath, []byte(data), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func WriteBytesLn(rawPath string, data []byte) {
	tmp := make([]byte, 0, len(data)+1)
	tmp = append(tmp, data...)
	tmp = append(tmp, '\n')
	err := os.WriteFile(rawPath, tmp, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
func WriteStringLn(rawPath string, data string) {
	err := os.WriteFile(rawPath, []byte(data+"\n"), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func AppendBytes(rawPath string, data []byte) int {
	file, err := os.OpenFile(rawPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	writeLen, err := file.Write(data)
	if err != nil {
		panic(err)
	}
	return writeLen
}
func AppendString(rawPath string, data string) int {
	file, err := os.OpenFile(rawPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	writeLen, err := file.WriteString(data)
	if err != nil {
		panic(err)
	}
	return writeLen
}

func AppendBytesLn(rawPath string, data []byte) int {
	file, err := os.OpenFile(rawPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	writeLen, err := file.Write(data)
	if err != nil {
		panic(err)
	}
	_, err = file.Write([]byte{'\n'})
	if err != nil {
		panic(err)
	}
	return writeLen
}
func AppendStringLn(rawPath string, data string) int {
	file, err := os.OpenFile(rawPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	writeLen, err := file.WriteString(data)
	if err != nil {
		panic(err)
	}
	_, err = file.Write([]byte{'\n'})
	if err != nil {
		panic(err)
	}
	return writeLen
}

type AppendFile struct {
	raw *os.File
}

func OpenAppendFile(rawPath string) *AppendFile {
	file, err := os.OpenFile(rawPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return &AppendFile{raw: file}
}

func (receiver *AppendFile) Write(data []byte) int {
	writeLen, err := receiver.raw.Write(data)
	if err != nil {
		panic(err)
	}
	return writeLen
}
func (receiver *AppendFile) WriteString(data string) int {
	writeLen, err := receiver.raw.WriteString(data)
	if err != nil {
		panic(err)
	}
	return writeLen
}

func (receiver *AppendFile) WriteLn(data []byte) int {
	writeLen, err := receiver.raw.Write(data)
	if err != nil {
		panic(err)
	}
	receiver.WriteLnChar()
	return writeLen
}
func (receiver *AppendFile) WriteStringLn(data string) int {
	writeLen, err := receiver.raw.WriteString(data)
	if err != nil {
		panic(err)
	}
	receiver.WriteLnChar()
	return writeLen
}

func (receiver *AppendFile) WriteLnChar() {
	_, err := receiver.raw.Write([]byte{'\n'})
	if err != nil {
		panic(err)
	}
}

func (receiver *AppendFile) Close() {
	err := receiver.raw.Close()
	if err != nil {
		panic(err)
	}
}
