package carbon

import (
	"strings"
	"time"
)

// format转layout
func format2layout(format string) string {
	format = strings.Trim(format, " ")
	layout := strings.Replace(format, "Y", "2006", 1)
	layout = strings.Replace(layout, "y", "06", 1)
	layout = strings.Replace(layout, "m", "01", 1)
	layout = strings.Replace(layout, "n", "1", 1)
	layout = strings.Replace(layout, "d", "02", 1)
	layout = strings.Replace(layout, "j", "2", 1)
	layout = strings.Replace(layout, "H", "15", 1)
	layout = strings.Replace(layout, "h", "03", 1)
	layout = strings.Replace(layout, "g", "3", 1)
	layout = strings.Replace(layout, "i", "04", 1)
	layout = strings.Replace(layout, "s", "05", 1)
	return layout
}

// getStartDay 获取开始日期
func (c *Carbon) getStartDay(t time.Time) *Carbon {
	return c.CreateFromDate(t.Year(), t.Month(), t.Day())
}

// getEndDay 获取结束日期
func (c *Carbon) getEndDay(t time.Time) *Carbon {
	days := 30
	if c.Time.Month() == time.February {
		days = 28
	}
	if c.IsLeapYear() {
		days = days + 1
	}
	return c.CreateFromDateTime(t.Year(), t.Month(), days-1, HoursPerDay-1, MinutesPerHour-1, SecondsPerMinute-1)
}
