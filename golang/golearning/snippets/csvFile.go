package main

import (
    "os"
    "log"
    "encoding/csv"
)

var data = [][]string{{"Line1", "Hello Readers of:"}, {"Line2", "golangcode.com"}}

func main() {
    file, err := os.Create("result.csv")
    checkError("Cannot create file", err)
    defer file.Close()

    writer := csv.NewWriter(file)

    for _, value := range data {
        err := writer.Write(value)
        checkError("Cannot write to file", err)
    }

    defer writer.Flush()
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

func csvFileReader() {
    f, _ := os.Open("C:\\programs\\file.txt")
    defer f.Close()

    r := csv.NewReader(bufio.NewReader(f))

    for {
        record, err := r.Read()

        if err == io.EOF {
            break
        }

        fmt.Println(record)
        fmt.Println(len(record))

        for value := range record {
            fmt.Printf("  %v\n", record[value])
        }
    }
}
