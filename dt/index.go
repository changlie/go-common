package dt

import "time"

const DTFORMAT = "2006-01-02 15:04:05"
const DFORMAT = "2006-01-02"

type DateTime struct {
	raw time.Time
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
func (receiver *DateTime) Format(layout string) string {
	return receiver.raw.Format(layout)
}
func (receiver *DateTime) String() string {
	return receiver.raw.String()
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
