package carbon

import (
	"strings"
	"time"
)

type Carbon struct {
	Time time.Time
	loc  *time.Location
}

// New 初始化,设置默认时区
func New() *Carbon {
	loc, _ := time.LoadLocation(Local)
	return &Carbon{loc: loc}
}

// Timezone 设置时区，必须在New()之后，其他方法之前调用
func (c *Carbon) Timezone(name string) *Carbon {
	loc, _ := time.LoadLocation(name)
	c.loc = loc
	return c
}

// CreateFromTimestamp 时间戳转时间类型
func (c *Carbon) CreateFromTimestamp(timestamp int64) *Carbon {
	c.Time = time.Unix(timestamp, 0)
	return c
}

// CreateFromDateTime 年月日时分秒转时间类型
func (c *Carbon) CreateFromDateTime(year int, month time.Month, day int, hour int, minute int, second int) *Carbon {
	c.Time = time.Date(year, month, day, hour, minute, second, 0, c.loc)
	return c
}

// CreateFromDate 年月日转时间类型
func (c *Carbon) CreateFromDate(year int, month time.Month, day int) *Carbon {
	c.Time = time.Date(year, month, day, 0, 0, 0, 0, c.loc)
	return c
}

// CreateFromTime 时分秒转时间类型
func (c *Carbon) CreateFromTime(hour int, minute int, second int) *Carbon {
	c.Time = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, minute, second, 0, c.loc)
	return c
}

// Now 当前
func (c *Carbon) Now() *Carbon {
	c.Time = time.Now()
	return c
}

// AddYears N年后
func (c *Carbon) AddYears(years int) *Carbon {
	c.Time = c.Time.AddDate(years, 0, 0)
	return c
}

// AddYear 1年后
func (c *Carbon) AddYear() *Carbon {
	c.Time = c.Time.AddDate(1, 0, 0)
	return c
}

// SubYears N年前
func (c *Carbon) SubYears(years int) *Carbon {
	c.Time = c.Time.AddDate(-years, 0, 0)
	return c
}

// SubYear 1年前
func (c *Carbon) SubYear() *Carbon {
	c.Time = c.Time.AddDate(-1, 0, 0)
	return c
}

// AddMonths N月后
func (c *Carbon) AddMonths(months int) *Carbon {
	c.Time = c.Time.AddDate(0, months, 0)
	return c
}

// AddMonth 1月后
func (c *Carbon) AddMonth() *Carbon {
	c.Time = c.Time.AddDate(0, 1, 0)
	return c
}

// SubMonths N月前
func (c *Carbon) SubMonths(months int) *Carbon {
	c.Time = c.Time.AddDate(0, -months, 0)
	return c
}

// SubMonth 1月前
func (c *Carbon) SubMonth() *Carbon {
	c.Time = c.Time.AddDate(0, -1, 0)
	return c
}

// AddWeeks N周后
func (c *Carbon) AddWeeks(weeks int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, weeks*DaysPerWeek)
	return c
}

// AddWeek 1周后
func (c *Carbon) AddWeek() *Carbon {
	c.Time = c.Time.AddDate(0, 0, DaysPerWeek)
	return c
}

// SubWeeks N周前
func (c *Carbon) SubWeeks(weeks int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -weeks*DaysPerWeek)
	return c
}

// SubMonth 1月前
func (c *Carbon) SubWeek() *Carbon {
	c.Time = c.Time.AddDate(0, 0, -DaysPerWeek)
	return c
}

// AddDays N天后
func (c *Carbon) AddDays(days int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, days)
	return c
}

// AddDay 1天后
func (c *Carbon) AddDay() *Carbon {
	c.Time = c.Time.AddDate(0, 0, 1)
	return c
}

// SubDays N天前
func (c *Carbon) SubDays(days int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -days)
	return c
}

// SubDay 1天前
func (c *Carbon) SubDay() *Carbon {
	c.Time = c.Time.AddDate(0, 0, -1)
	return c
}

// AddHours N小时后
func (c *Carbon) AddHours(hours int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, hours/HoursPerDay)
	return c
}

// AddHour 1小时后
func (c *Carbon) AddHour() *Carbon {
	c.Time = c.Time.AddDate(0, 0, 1/HoursPerDay)
	return c
}

// AddHours N小时前
func (c *Carbon) SubHours(hours int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -hours/HoursPerDay)
	return c
}

// AddHour 1小时前
func (c *Carbon) SubHour() *Carbon {
	c.Time = c.Time.AddDate(0, 0, -1/HoursPerDay)
	return c
}

// AddMinutes N分钟后
func (c *Carbon) AddMinutes(minutes int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, minutes/MinutesPerDay)
	return c
}

// AddMinute 1分钟后
func (c *Carbon) AddMinute() *Carbon {
	c.Time = c.Time.AddDate(0, 0, 1/MinutesPerDay)
	return c
}

// SubMinutes N分钟前
func (c *Carbon) SubMinutes(minutes int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -minutes/MinutesPerDay)
	return c
}

// SubMinute 1分钟前
func (c *Carbon) SubMinute() *Carbon {
	c.Time = c.Time.AddDate(0, 0, -1/MinutesPerDay)
	return c
}

// AddSeconds N秒钟后
func (c *Carbon) AddSeconds(second int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, second/SecondsPerDay)
	return c
}

// AddMinute 1秒钟后
func (c *Carbon) AddSecond() *Carbon {
	c.Time = c.Time.AddDate(0, 0, 1/SecondsPerDay)
	return c
}

// SubMinutes N秒钟前
func (c *Carbon) SubSeconds(second int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -second/SecondsPerDay)
	return c
}

// SubMinute 1秒钟前
func (c *Carbon) SubSecond() *Carbon {
	c.Time = c.Time.AddDate(0, 0, -1/SecondsPerDay)
	return c
}

// Parse 解析标准时间格式
func (c *Carbon) Parse(value string) *Carbon {
	value = strings.Trim(value, " ")
	layout := "2006-01-02 15:04:05"

	if len(value) == 10 && strings.Contains(value, "-") {
		layout = "2006-01-02"
	}

	if len(value) == 10 && strings.Contains(value, "/") {
		layout = "2006/01/02"
	}

	if len(value) == 19 && strings.Contains(value, "/") {
		layout = "2006/01/02 15:04:05"
	}

	if len(value) == 14 {
		layout = "20060102150405"
	}

	if len(value) == 8 {
		layout = "20060102"
	}

	t, _ := time.ParseInLocation(layout, value, c.loc)
	c.Time = t
	return c
}

// ParseByCustom 解析自定义时间格式
func (c *Carbon) ParseByCustom(value string, format string) *Carbon {
	value = strings.Trim(value, " ")
	layout := format2layout(format)
	t, _ := time.ParseInLocation(layout, value, c.loc)
	c.Time = t
	return c
}

// Format 格式化时间
func (c *Carbon) Format(format string) string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.loc).Format(format2layout(format))
}

// Today 今天
func (c *Carbon) Today() string {
	now := time.Now()
	return now.In(c.loc).Format("2006-01-02")
}

// StartOfToday 今天开始时间
func (c *Carbon) StartOfToday() string {
	now := time.Now()
	return c.getStartDay(now).ToDateTimeString()
}

// EndOfToday 今天结束时间
func (c *Carbon) EndOfToday() string {
	now := time.Now()
	return c.getEndDay(now).ToDateTimeString()
}

// Tomorrow 明天
func (c *Carbon) Tomorrow() string {
	tomorrow := time.Now().AddDate(0, 0, 1)
	return tomorrow.In(c.loc).Format("2006-01-02")
}

// StartOfTomorrow 明天开始时间
func (c *Carbon) StartOfTomorrow() string {
	tomorrow := time.Now().AddDate(0, 0, 1)
	return c.getStartDay(tomorrow).ToDateTimeString()
}

// EndOfTomorrow 明天结束时间
func (c *Carbon) EndOfTomorrow() string {
	tomorrow := time.Now().AddDate(0, 0, 1)
	return c.getEndDay(tomorrow).ToDateTimeString()
}

// Yesterday 昨天
func (c *Carbon) Yesterday() string {
	yesterday := time.Now().AddDate(0, 0, -1)
	return yesterday.In(c.loc).Format("2006-01-02")
}

// StartOfYesterday 昨天开始时间
func (c *Carbon) StartOfYesterday() string {
	yesterday := time.Now().AddDate(0, 0, -1)
	return c.getStartDay(yesterday).ToDateTimeString()
}

// EndOfYesterday 昨天结束时间
func (c *Carbon) EndOfYesterday() string {
	yesterday := time.Now().AddDate(0, 0, -1)
	return c.getEndDay(yesterday).ToDateTimeString()
}

// FirstDayInYear 年初第一天
func (c *Carbon) FirstOfYear() string {
	return c.CreateFromDate(c.Time.Year(), 01, 01).ToDateStartString()
}

// LastDayInYear 年末最后一天
func (c *Carbon) LastOfYear() string {
	return c.CreateFromDate(c.Time.Year(), 12, 31).ToDateStartString()
}

// FirstDayInMonth 月初第一天
func (c *Carbon) FirstOfMonth() string {
	return c.CreateFromDate(c.Time.Year(), c.Time.Month(), 01).ToDateStartString()
}

// LastDayInMonth 月末最后一天
func (c *Carbon) LastOfMonth() string {
	return c.getEndDay(c.Time).ToDateStartString()
}

// StartOfYear 当年开始时间
func (c *Carbon) StartOfYear() string {
	return c.CreateFromDate(c.Time.Year(), 01, 01).Format("2006-01-02 00:00:00")
}

// EndOfYear 当年结束时间
func (c *Carbon) EndOfYear() string {
	return c.CreateFromDateTime(c.Time.Year(), 12, 31, HoursPerDay-1, MinutesPerHour-1, SecondsPerMinute-1).Format("2006-01-02 15:04:05")
}

// StartOfMonth 当月开始时间
func (c *Carbon) StartOfMonth() string {
	t := c.CreateFromDate(c.Time.Year(), c.Time.Month(), c.Time.Day()).Time
	return c.getStartDay(t).ToDateTimeString()
}

// EndOfMonth 当月结束时间
func (c *Carbon) EndOfMonth() string {
	t := c.CreateFromDate(c.Time.Year(), c.Time.Month(), c.Time.Day()).Time
	return c.getEndDay(t).ToDateTimeString()
}

// FirstDayInMonth 当天开始时间
func (c *Carbon) StartOfDay() string {
	return c.getStartDay(c.Time).ToDateTimeString()
}

// LastDayInMonth 当天结束时间
func (c *Carbon) EndOfDay() string {
	return c.getEndDay(c.Time).ToDateTimeString()
}

// ToDateTimeString 转日期时间字符串
func (c *Carbon) ToDateTimeString() string {
	return c.Time.In(c.loc).Format("2006-01-02 15:04:05")
}

// ToDateString 转日期字符串
func (c *Carbon) ToDateString() string {
	return c.Time.In(c.loc).Format("2006-01-02")
}

// ToDateStartString 转日期开始时间字符串
func (c *Carbon) ToDateStartString() string {
	return c.Time.In(c.loc).Format("2006-01-02 00:00:00")
}

// ToDateEndString 转日期结束时间字符串
func (c *Carbon) ToDateEndString() string {
	return c.CreateFromDateTime(c.Time.Year(), c.Time.Month(), c.Time.Day(), HoursPerDay-1, MinutesPerHour-1, SecondsPerMinute-1).Format("2006-01-02 15:04:05")
}

// ToTimeString 转时间字符串
func (c *Carbon) ToTimeString() string {
	return c.Time.In(c.loc).Format("15:04:05")
}

// ToTimeStartString 转开始时间字符串
func (c *Carbon) ToTimeStartString() string {
	return c.Time.In(c.loc).Format("15:00:00")
}

// ToDateEndString 转结束时间字符串
func (c *Carbon) ToTimeEndString() string {
	return c.CreateFromDateTime(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), MinutesPerHour-1, SecondsPerMinute-1).Format("15:04:05")
}

// ToTimestamp 转时间戳
func (c *Carbon) ToTimestamp() int64 {
	return c.Time.Unix()
}

// IsLeapYear 是否是闰年
func (c *Carbon) IsLeapYear() bool {
	year := c.Time.Year()
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// IsJanuary 是否是一月
func (c *Carbon) IsJanuary() bool {
	return c.Time.In(c.loc).Month() == time.January
}

// IsMonday 是否是二月
func (c *Carbon) IsFebruary() bool {
	return c.Time.In(c.loc).Month() == time.February
}

// IsMarch 是否是三月
func (c *Carbon) IsMarch() bool {
	return c.Time.In(c.loc).Month() == time.March
}

// IsApril 是否是四月
func (c *Carbon) IsApril() bool {
	return c.Time.In(c.loc).Month() == time.April
}

// IsMay 是否是五月
func (c *Carbon) IsMay() bool {
	return c.Time.In(c.loc).Month() == time.May
}

// IsJune 是否是六月
func (c *Carbon) IsJune() bool {
	return c.Time.In(c.loc).Month() == time.June
}

// IsJuly 是否是七月
func (c *Carbon) IsJuly() bool {
	return c.Time.In(c.loc).Month() == time.July
}

// IsAugust 是否是八月
func (c *Carbon) IsAugust() bool {
	return c.Time.In(c.loc).Month() == time.August
}

// IsSeptember 是否是九月
func (c *Carbon) IsSeptember() bool {
	return c.Time.In(c.loc).Month() == time.September
}

// IsOctober 是否是十月
func (c *Carbon) IsOctober() bool {
	return c.Time.In(c.loc).Month() == time.October
}

// IsNovember 是否是十一月
func (c *Carbon) IsNovember() bool {
	return c.Time.In(c.loc).Month() == time.November
}

// IsDecember 是否是十二月
func (c *Carbon) IsDecember() bool {
	return c.Time.In(c.loc).Month() == time.December
}

// IsMonday 是否是周一
func (c *Carbon) IsMonday() bool {
	return c.Time.In(c.loc).Weekday() == time.Monday
}

// IsTuesday 是否是周二
func (c *Carbon) IsTuesday() bool {
	return c.Time.In(c.loc).Weekday() == time.Tuesday
}

// IsWednesday 是否是周三
func (c *Carbon) IsWednesday() bool {
	return c.Time.In(c.loc).Weekday() == time.Wednesday
}

// IsThursday 是否是周四
func (c *Carbon) IsThursday() bool {
	return c.Time.In(c.loc).Weekday() == time.Thursday
}

// IsFriday 是否是周五
func (c *Carbon) IsFriday() bool {
	return c.Time.In(c.loc).Weekday() == time.Friday
}

// IsSaturday 是否是周六
func (c *Carbon) IsSaturday() bool {
	return c.Time.In(c.loc).Weekday() == time.Saturday
}

// IsSunday 是否是周日
func (c *Carbon) IsSunday() bool {
	return c.Time.In(c.loc).Weekday() == time.Sunday
}

// IsToday 是否是今天
func (c *Carbon) IsToday() bool {
	return c.ToDateString() == c.Today()
}

// IsYesterday 是否是昨天
func (c *Carbon) IsYesterday() bool {
	return c.SubDay().ToDateString() == c.Yesterday()
}

// IsTomorrow 是否是明天
func (c *Carbon) IsTomorrow() bool {
	return c.AddDay().ToDateString() == c.Tomorrow()
}

// IsFirstOfYear 是否年初
func (c *Carbon) IsFirstOfYear() bool {
	_, month, day := c.Time.Date()
	if month == time.January && day == 1 {
		return true
	}
	return false
}

// IsLastOfYear 是否是年末
func (c *Carbon) IsLastOfYear() bool {
	_, month, day := c.Time.Date()
	if month == time.December && day == 31 {
		return true
	}
	return false
}

// IsFirstOfMonth 是否月初
func (c *Carbon) IsFirstOfMonth() bool {
	return c.ToDateTimeString() == c.FirstOfMonth()
}

// IsLastOfMonth 是否是月末
func (c *Carbon) IsLastOfMonth() bool {
	return c.ToDateTimeString() == c.LastOfMonth()
}

// IsStartOfYear 是否当年开始时间
func (c *Carbon) IsStartOfYear() bool {
	return c.ToDateTimeString() == c.StartOfYear()
}

// IsEndOfYear 是否当年结束时间
func (c *Carbon) IsEndOfYear() bool {
	return c.ToDateTimeString() == c.EndOfYear()
}

// IsStartOfMonth 是否当月开始时间
func (c *Carbon) IsStartOfMonth() bool {
	return c.ToDateTimeString() == c.StartOfMonth()
}

// IsEndOfMonth 是否当月结束时间
func (c *Carbon) IsEndOfMonth() bool {
	return c.ToDateTimeString() == c.EndOfMonth()
}

// IsStartOfDay 是否当天开始时间
func (c *Carbon) IsStartOfDay() bool {
	return c.ToDateTimeString() == c.StartOfDay()
}

// IsEndOfDay 是否当天结束时间
func (c *Carbon) IsEndOfDay() bool {
	return c.ToDateTimeString() == c.EndOfDay()
}

// IsStartOfToday 是否今天开始时间
func (c *Carbon) IsStartOfToday() bool {
	return c.ToDateTimeString() == c.StartOfToday()
}

// IsEndOfToday 是否今天结束时间
func (c *Carbon) IsEndOfToday() bool {
	return c.ToDateTimeString() == c.EndOfToday()
}

// IsStartOfTomorrow 是否明天开始时间
func (c *Carbon) IsStartOfTomorrow() bool {
	return c.ToDateTimeString() == c.StartOfTomorrow()
}

// IsEndOfTomorrow 是否明天结束时间
func (c *Carbon) IsEndOfTomorrow() bool {
	return c.ToDateTimeString() == c.EndOfTomorrow()
}

// IsStartOfYesterday 是否昨天开始时间
func (c *Carbon) IsStartOfYesterday() bool {
	return c.ToDateTimeString() == c.StartOfYesterday()
}

// IsEndOfYesterday 是否昨天结束时间
func (c *Carbon) IsEndOfYesterday() bool {
	return c.ToDateTimeString() == c.EndOfYesterday()
}
