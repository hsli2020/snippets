package mdate

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

const dateStrFormat = "2006-01-02"

type Date struct {
	time.Time
}

type Dates []Date

func NewDateFromStrWithFormat(format string, str string) (Date, error) {
	t, err := time.Parse(format, str)
	if err != nil {
		return Date{}, err
	}
	return Date{t}, nil
}

func NewDate(year int, month time.Month, day int) Date {
	return Date{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

func GetToday() Date {
	y, m, d := time.Now().Date()
	return Date{time.Date(y, m, d, 0, 0, 0, 0, time.UTC)}
}

func NewDateFromStr(str string) (Date, error) {
	return NewDateFromStrWithFormat(dateStrFormat, str)
}

func MustDateFromStr(str string) Date {
	date, err := NewDateFromStr(str)
	if err != nil {
		panic(err)
	}
	return date
}

func (d *Date) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("\"2006-01-02\"", string(data))
	*d = Date{t}
	return err
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d Date) PlusNDay(n int) Date {
	return Date{d.AddDate(0, 0, n)}
}

func (d Date) MinusNDay(n int) Date {
	return Date{d.AddDate(0, 0, -n)}
}

func (d Date) IsEqual(other Date) bool {
	y1, m1, d1 := d.Date()
	y2, m2, d2 := other.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func (d Date) IsEarlier(other Date) bool {
	return d.Unix() < other.Unix()
}

func (d Date) IsEarlierEq(other Date) bool {
	return d.Unix() <= other.Unix()
}

func (d Date) IsLater(other Date) bool {
	return d.Unix() > other.Unix()
}

func (d Date) IsLaterEq(other Date) bool {
	return d.Unix() >= other.Unix()
}

func (d Date) IsNextDay(other Date) bool {
	otherPlus1Day := other.PlusNDay(1)
	return d.IsEqual(otherPlus1Day)
}

func (d Date) DateDiff(other Date) int {
	duration := other.Sub(d.Time)
	return int(duration.Hours() / 24)
}

func (d Date) ToStringWithFormat(format string) string {
	return d.Format(format)
}

func (d Date) String() string {
	return d.ToStringWithFormat(dateStrFormat)
}

// for sql driver
func (d *Date) Scan(value interface{}) error {
	d.Time = value.(time.Time)
	return nil
}

// for sql driver
func (d Date) Value() (driver.Value, error) {
	return d.Time, nil
}

func (dates Dates) ToStringArray() []string {
	ret := []string{}
	for _, date := range dates {
		ret = append(ret, date.String())
	}
	return ret
}

func (d Dates) Len() int {
	return len(d)
}

func (d Dates) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Dates) Less(i, j int) bool {
	return d[i].IsEarlierEq(d[j])
}

func (d Date) JapanFiscalYear() JapanFiscalYear {
	if d.Month() <= 3 {
		return JapanFiscalYear(d.Year() - 1)
	}
	return JapanFiscalYear(d.Year())
}
