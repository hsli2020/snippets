package main

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "USER:PASSWORD@/DATABASE")
    defer db.Close()

    if err != nil {
        log.Fatal(err.Error())
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err.Error())
    }


    rows, err := db.Query("SELECT COUNT(*) AS total FROM wp_posts")
    defer rows.Close()

    if err != nil {
        log.Fatal(err.Error())
    }

    for rows.Next() {
        var total int
        if err := rows.Scan(&total); err != nil {
                log.Fatal(err.Error())
        }
        fmt.Printf("Total rows found = %d\n", total)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err.Error())
    }
}
