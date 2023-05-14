package dt

import (
	"strings"
	"time"
)

const DTFORMAT = "2006-01-02 15:04:05"
const DFORMAT = "2006-01-02"

type DateTime struct {
	raw time.Time
}

func Fmt(raw string) string {
	var tmp string
	tmp = strings.Replace(raw, "y", "2006", 1)
	tmp = strings.Replace(tmp, "M", "01", 1)
	tmp = strings.Replace(tmp, "d", "02", 1)
	tmp = strings.Replace(tmp, "h", "15", 1)
	tmp = strings.Replace(tmp, "m", "04", 1)
	tmp = strings.Replace(tmp, "s", "05", 1)
	return tmp
}

func Now() *DateTime {
	return &DateTime{raw: time.Now()}
}

func Parse(layout, value string) *DateTime {
	tmp, err := time.Parse(layout, value)
	if err != nil {
		return nil
	}
	return &DateTime{raw: tmp}
}

func (receiver *DateTime) AddYear(years int) *DateTime {
	return &DateTime{raw: receiver.raw.AddDate(years, 0, 0)}
}
func (receiver *DateTime) AddMonth(months int) *DateTime {
	return &DateTime{raw: receiver.raw.AddDate(0, months, 0)}
}
func (receiver *DateTime) AddDay(days int) *DateTime {
	return &DateTime{raw: receiver.raw.AddDate(0, 0, days)}
}
func (receiver *DateTime) Add(d time.Duration) *DateTime {
	return &DateTime{raw: receiver.raw.Add(d)}
}

func (receiver *DateTime) Format(layout string) string {
	return receiver.raw.Format(layout)
}

func (receiver *DateTime) String() string {
	return receiver.raw.String()
}

func (receiver *DateTime) ResetDay() *DateTime {
	return &DateTime{raw: receiver.raw.AddDate(0, 0, -receiver.raw.Day()+1)}
}
func (receiver *DateTime) ResetHour() *DateTime {
	return &DateTime{raw: receiver.raw.Add(-time.Duration(receiver.raw.Hour()) * time.Hour)}
}
func (receiver *DateTime) ResetMinute() *DateTime {
	return &DateTime{raw: receiver.raw.Add(-time.Duration(receiver.raw.Minute()) * time.Minute)}
}
func (receiver *DateTime) ResetSecond() *DateTime {
	return &DateTime{raw: receiver.raw.Add(-time.Duration(receiver.raw.Second()) * time.Second)}
}

func (receiver *DateTime) SetYear(years int) *DateTime {
	return &DateTime{raw: receiver.raw.AddDate(-receiver.raw.Year()+years, 0, 0)}
}
func (receiver *DateTime) SetMonth(months int) *DateTime {
	return &DateTime{raw: receiver.raw.AddDate(0, -int(receiver.raw.Month())+months, 0)}
}
func (receiver *DateTime) SetDay(days int) *DateTime {
	return &DateTime{raw: receiver.raw.AddDate(0, 0, -receiver.raw.Day()+days)}
}
func (receiver *DateTime) SetHour(hours int) *DateTime {
	return &DateTime{raw: receiver.raw.Add(-time.Duration(receiver.raw.Hour())*time.Hour + time.Duration(hours)*time.Hour)}
}
func (receiver *DateTime) SetMinute(minutes int) *DateTime {
	return &DateTime{raw: receiver.raw.Add(-time.Duration(receiver.raw.Hour())*time.Hour + time.Duration(minutes)*time.Hour)}
}
func (receiver *DateTime) SetSecond(seconds int) *DateTime {
	return &DateTime{raw: receiver.raw.Add(-time.Duration(receiver.raw.Second())*time.Second + time.Duration(seconds)*time.Second)}
}

func (receiver *DateTime) Year() int {
	return receiver.raw.Year()
}
func (receiver *DateTime) Month() int {
	return int(receiver.raw.Month())
}
func (receiver *DateTime) Day() int {
	return receiver.raw.Day()
}
func (receiver *DateTime) Hour() int {
	return receiver.raw.Hour()
}
func (receiver *DateTime) Minute() int {
	return receiver.raw.Minute()
}
func (receiver *DateTime) Second() int {
	return receiver.raw.Second()
}

func (receiver *DateTime) Weekday() int {
	return int(receiver.raw.Weekday())
}