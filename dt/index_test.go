package dt

import (
	"fmt"
	"testing"
	"time"
)

func Test_d2(t *testing.T) {
	fmt.Println(time.Now().Format(DTFORMAT))
	fmt.Println(time.Now().Add(time.Minute * 5).Format(DTFORMAT))
	fmt.Println(time.Now().Add(time.Minute * -17).Format(DTFORMAT))
}
func Test_d1(t *testing.T) {
	fmt.Println(time.Now().Format(DTFORMAT))
	fmt.Println(time.Now().Format(DFORMAT))
	fmt.Println(time.Now().AddDate(0, -6, 0).Format(DTFORMAT))

	fmt.Println(time.Parse(DFORMAT, "1993-07-23"))

	res := Parse(DFORMAT, "1993-07-23")
	fmt.Println(res)
	fmt.Println(res.Format(DTFORMAT))
	fmt.Println("res.Month()", res.Month())
	fmt.Println(res.AddDay(22).Month())
	fmt.Println(res.AddMonth(6).Month())
	fmt.Println(res.AddMonth(6).Format(DFORMAT))
	fmt.Println(Now().Month())
	fmt.Println("Now().Weekday()", Now().Weekday())
	fmt.Println("Now().Weekday()", Now().AddDay(1).Weekday())
	fmt.Println("Now().Weekday()", Now().AddDay(2).Weekday())

}
