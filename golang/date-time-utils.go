func GetTimeStamp() string {
     loc, _ := time.LoadLocation("America/Los_Angeles")
     t := time.Now().In(loc)
     return t.Format("20060102150405")
}
func GetTodaysDate() string {
    loc, _ := time.LoadLocation("America/Los_Angeles")
    current_time := time.Now().In(loc)
    return current_time.Format("2006-01-02")
}

func GetTodaysDateTime() string {
    loc, _ := time.LoadLocation("America/Los_Angeles")
    current_time := time.Now().In(loc)
    return current_time.Format("2006-01-02 15:04:05")
}

func GetTodaysDateTimeFormatted() string {
    loc, _ := time.LoadLocation("America/Los_Angeles")
    current_time := time.Now().In(loc)
    return current_time.Format("Jan 2, 2006 at 3:04 PM")
}

func GetTimeStampFromDate(dtformat string) string {
    form := "Jan 2, 2006 at 3:04 PM"
    t2, _ := time.Parse(form, dtformat)
    return t2.Format("20060102150405")
}

// Parse unix timestamp string
package main

import (
    "fmt"
    "time"
    "strconv"
)

func main() {
    i, err := strconv.ParseInt("1405544146", 10, 64)
    if err != nil {
        panic(err)
    }
    tm := time.Unix(i, 0)
    fmt.Println(tm)
}