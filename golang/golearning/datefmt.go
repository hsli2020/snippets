package main

import (
    "fmt"
    "strings"
    "time"
)

func FmtDateTime(format string, t time.Time) string {
    // Y-m-d H:i:s => 2006-01-02 15:04:05
    r := strings.NewReplacer(
        "Y", "2006",
        "m", "01",
        "d", "02",
        "H", "15",
        "i", "04",
        "s", "05",
    )
    format = r.Replace(format)
    return t.Format(format)
}

func main() {
    fmt.Println(FmtDateTime("Y-m-d H:i:s", time.Now()))
}
