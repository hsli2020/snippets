package main

import (
    "fmt"               // builtin package
    "time"              // builtin package

    "stringutil"        // project package
    "utils/codegen"     // project package
    "dateutil"          // project package
    "fileutil"          // project package
    "dbutil"            // project package

    "github.com/leekchan/timeutil"
)

func main() {
    // stringutil
    fmt.Println(stringutil.Reverse("!oG ,olleH"))

    // dateutil
    fmt.Println(dateutil.FmtDateTime("Y-m-d H:i:s", time.Now()))

    // codegen
    strfmt := fmt.Sprintf

    lines := codegen.NewCodeLines()

    lines.Push("<Request>")
    lines.Push("<Item>")
    lines.Push(strfmt("<Weight>%d</Weight>", 23))
    lines.Push("<Height>%d</Height>", 45)
    lines.Push("</Item>")
    lines.Push("</Request>")

    fmt.Println(lines.ToString())

    // fileutil
    fmt.Println(fileutil.Exists("./main.go"))
    fmt.Println(fileutil.Exists("./main.exe"))

    // dbutil
    columns := []string{"sku", "qty", "price"}

    items := make([]map[string]string, 0)

    for i:=1; i<10; i++ {
        items = append(items, map[string]string{
            "sku":   fmt.Sprintf("SKU-ABC-%d", i),
            "qty":   fmt.Sprintf("%d", i),
            "price": fmt.Sprintf("%d.%d", i*11, i*11),
        })
    }

    fmt.Println(dbutil.InsertSql("mytable", columns, items));

	// github.com/leekchan/timeutil	
	date := time.Date(2015, 7, 2, 15, 24, 30, 35, time.UTC)
	fmt.Println(timeutil.Strftime(&date, "%U %W"))
}
