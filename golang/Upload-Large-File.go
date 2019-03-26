//NOTE: for simplicity, error check is omitted
func uploadLargeFile(uri, filePath string, chunkSize int, params map[string]string) {
    //open file and retrieve info
    file, _ := os.Open(filePath)
    fi, _ := file.Stat()
    defer file.Close()

    //buffer for storing multipart data
    byteBuf := &bytes.Buffer{}

    //part: parameters
    mpWriter := multipart.NewWriter(byteBuf)
    for key, value := range params {
        _ = mpWriter.WriteField(key, value)
    }

    //part: file
    mpWriter.CreateFormFile("file", fi.Name())
    contentType := mpWriter.FormDataContentType()

    nmulti := byteBuf.Len()
    multi := make([]byte, nmulti)
    _, _ = byteBuf.Read(multi)

    //part: latest boundary
    //when multipart closed, latest boundary is added
    mpWriter.Close()
    nboundary := byteBuf.Len()
    lastBoundary := make([]byte, nboundary)
    _, _ = byteBuf.Read(lastBoundary)

    //calculate content length
    totalSize := int64(nmulti) + fi.Size() + int64(nboundary)
    log.Printf("Content length = %v byte(s)\n", totalSize)

    //use pipe to pass request
    rd, wr := io.Pipe()
    defer rd.Close()

    go func() {
        defer wr.Close()

        //write multipart
        _, _ = wr.Write(multi)

        //write file
        buf := make([]byte, chunkSize)
        for {
            n, err := file.Read(buf)
            if err != nil {
                break
            }
            _, _ = wr.Write(buf[:n])
        }
        //write boundary
        _, _ = wr.Write(lastBoundary)
    }()

    //construct request with rd
    req, _ := http.NewRequest("POST", uri, rd)
    req.Header.Set("Content-Type", contentType)
    req.ContentLength = totalSize

    //process request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    } else {
        log.Println(resp.StatusCode)
        log.Println(resp.Header)

        body := &bytes.Buffer{}
        _, _ = body.ReadFrom(resp.Body)
        resp.Body.Close()
        log.Println(body)
    }
}
