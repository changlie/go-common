package a

import "fmt"

func Exec(raw string) {
	fmt.Println("执行", raw)
}

func Echo(args ...any) {
	fmt.Println(args...)
}

func F(format string, a ...any) {
	fmt.Printf(format, a...)
}

func Fln(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}
