package a

import "testing"

func Test_uuid(t *testing.T) {
	s1 := Uuid()
	println(len(s1), s1)
	s2 := UuidV4()
	println(len(s2), s2)
}
