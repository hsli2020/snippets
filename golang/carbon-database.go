package carbon

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type ToDateTimeString time.Time
type ToDateString time.Time
type ToTimeString time.Time
type ToTimestamp time.Time

func (c *Carbon) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*c = Carbon{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
func (c Carbon) Value() (driver.Value, error) {
	var zeroTime time.Time
	var timeTime = c.Time
	if timeTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return timeTime, nil
}

func (t ToDateTimeString) MarshalJSON() ([]byte, error) {
	var timeFormat = ""
	var timeTime = time.Time(t)
	if !timeTime.IsZero() {
		timeFormat = timeTime.Format("2006-01-02 15:04:05")
	}
	return []byte(fmt.Sprintf("%s\"", timeFormat)), nil
}
func (t *ToDateTimeString) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = ToDateTimeString(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
func (t ToDateTimeString) Value() (driver.Value, error) {
	var zeroTime time.Time
	var timeTime = time.Time(t)
	if timeTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return timeTime, nil
}

func (t ToDateString) MarshalJSON() ([]byte, error) {
	var timeFormat = ""
	var timeTime = time.Time(t)
	if !timeTime.IsZero() {
		timeFormat = timeTime.Format("2006-01-02")
	}
	return []byte(fmt.Sprintf("\"%s\"", timeFormat)), nil
}
func (t *ToDateString) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = ToDateString(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
func (t ToDateString) Value() (driver.Value, error) {
	var zeroTime time.Time
	var timeTime = time.Time(t)
	if timeTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return timeTime, nil
}

func (t ToTimeString) MarshalJSON() ([]byte, error) {
	var timeFormat = ""
	var timeTime = time.Time(t)
	if !timeTime.IsZero() {
		timeFormat = timeTime.Format("15:04:05")
	}
	return []byte(fmt.Sprintf("\"%s\"", timeFormat)), nil
}
func (t *ToTimeString) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = ToTimeString(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
func (t ToTimeString) Value() (driver.Value, error) {
	var zeroTime time.Time
	var timeTime = time.Time(t)
	if timeTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return timeTime, nil
}

func (t ToTimestamp) MarshalJSON() ([]byte, error) {
	var timestamp = int64(0)
	var timeTime = time.Time(t)
	if !timeTime.IsZero() {
		timestamp = timeTime.Unix()
	}
	return []byte(fmt.Sprintf("%d", timestamp)), nil
}
func (t *ToTimestamp) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = ToTimestamp(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
func (t ToTimestamp) Value() (driver.Value, error) {
	var zeroTime time.Time
	var timeTime = time.Time(t)
	if timeTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return timeTime, nil
}
