
package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {

	var dates [4]time.Time

	dates[0], _ = time.Parse("2006-01-02 15:04:05.000000000 MST -07:00",
							 "1609-09-12 19:02:35.123456789 PDT +03:00")

	dates[1], _ = time.Parse("2006-01-02 03:04:05 PM -0700",
							 "1995-11-07 04:29:43 AM -0209")

	dates[2], _ = time.Parse("PM -0700 01/02/2006 03:04:05",
							 "AM -0209 11/07/1995 04:29:43")

	dates[3], _ = time.Parse("Time:Z07:00T15:04:05 Date:2006-01-02 ",
							 "Time:-03:30T19:18:35 Date:2119-10-29")

	defaultFormat := "2006-01-02 15:04:05 PM -07:00 Jan Mon MST"

	formats := []map[string]string{
		{"format": "2006", "description": "Year"},
		{"format": "06", "description": "Year"},

		{"format": "01", "description": "Month"},
		{"format": "1", "description": "Month"},
		{"format": "Jan", "description": "Month"},
		{"format": "January", "description": "Month"},

		{"format": "02", "description": "Day"},
		{"format": "2", "description": "Day"},

		{"format": "Mon", "description": "Week day"},
		{"format": "Monday", "description": "Week day"},

		{"format": "03", "description": "Hours"},
		{"format": "3", "description": "Hours"},
		{"format": "15", "description": "Hours"},

		{"format": "04", "description": "Minutes"},
		{"format": "4", "description": "Minutes"},

		{"format": "05", "description": "Seconds"},
		{"format": "5", "description": "Seconds"},

		{"format": "PM", "description": "AM or PM"},

		{"format": ".000", "description": "Miliseconds"},
		{"format": ".000000", "description": "Microseconds"},
		{"format": ".000000000", "description": "Nanoseconds"},

		{"format": "-0700", "description": "Timezone offset"},
		{"format": "-07:00", "description": "Timezone offset"},
		{"format": "Z0700", "description": "Timezone offset"},
		{"format": "Z07:00", "description": "Timezone offset"},

		{"format": "MST", "description": "Timezone"}}

	for _, date := range dates {
		fmt.Printf("\n\n %s \n", date.Format(defaultFormat))

		fmt.Printf("%-15s + %-12s + %12s \n",
			strings.Repeat("-", 15),
			strings.Repeat("-", 12),
			strings.Repeat("-", 12))

		fmt.Printf("%-15s | %-12s | %12s \n", "Type", "Placeholder", "Value")

		fmt.Printf("%-15s + %-12s + %12s \n",
			strings.Repeat("-", 15),
			strings.Repeat("-", 12),
			strings.Repeat("-", 12))

		for _, f := range formats {
			fmt.Printf("%-15s | %-12s | %-12s \n",
				f["description"],
				f["format"],
				date.Format(f["format"]))
		}

		fmt.Printf("%-15s + %-12s + %12s \n",
			strings.Repeat("-", 15),
			strings.Repeat("-", 12),
			strings.Repeat("-", 12))
	}
}
