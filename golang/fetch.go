import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func fetch (url string) string {
    fmt.Println("Fetch Url", url)
    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Http get err:", err)
        return ""
    }
    if resp.StatusCode != 200 {
        fmt.Println("Http status code:", resp.StatusCode)
        return ""
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Read error", err)
        return ""
    }
    return string(body)
}

/*

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    // http://httpbin.org/#/Anything/get_anything
    r, err := http.Get("http://httpbin.org/anything?hello=world")
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }
    fmt.Printf("body = %s\n", string(body))
}

*/