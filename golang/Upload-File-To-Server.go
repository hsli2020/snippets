func SendPostRequest (url string, filename string, filetype string) []byte {
    file, err := os.Open(filename)

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile(filetype, filepath.Base(file.Name()))

    if err != nil {
        log.Fatal(err)
    }

    io.Copy(part, file)
    writer.Close()
    request, err := http.NewRequest("POST", url, body)

    if err != nil {
        log.Fatal(err)
    }

    request.Header.Add("Content-Type", writer.FormDataContentType())
    client := &http.Client{}

    response, err := client.Do(request)

    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    content, err := ioutil.ReadAll(response.Body)

    if err != nil {
        log.Fatal(err)
    }

    return content
}
