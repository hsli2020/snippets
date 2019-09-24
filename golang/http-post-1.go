import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    url := "http://xxx/yyy"
    fmt.Println("URL:>", url)

    var query = []byte(`your query`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(query))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "text/plain")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}