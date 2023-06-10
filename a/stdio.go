package a

import "fmt"

func Echo(args ...any) {
	fmt.Println(args...)
}

func F(format string, a ...any) {
	fmt.Printf(format, a...)
}

func Fln(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}
