package main

import (
        "io/ioutil"
        "log"
        "net/http"
)

func main() {
        client := &http.Client{}

        req, err := http.NewRequest("GET", "http://httpbin.org/user-agent", nil)
        if err != nil {
                log.Fatalln(err)
        }

        req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

        resp, err := client.Do(req)
        if err != nil {
                log.Fatalln(err)
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                log.Fatalln(err)
        }

        log.Println(string(body))
}