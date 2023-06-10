package a

import (
	"github.com/fatih/color"
	"testing"
)

func Test_uuid(t *testing.T) {
	s1 := Uuid()
	println(len(s1), s1)
	s2 := UuidV4()
	println(len(s2), s2)
}

func TestConsoleColor(t *testing.T) {
	// Print with default helper functions
	color.Cyan("Prints text in cyan.")

	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// These are using the default foreground colors
	color.Red("We have red")
	color.Magenta("And many others ..")

}
