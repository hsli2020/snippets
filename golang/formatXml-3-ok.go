// https://stackoverflow.com/questions/21117161/go-how-would-you-pretty-print-prettify-html

import (
    "bytes"
    "encoding/xml"
    "io"
)

func formatXML(data []byte) ([]byte, error) {
    b := &bytes.Buffer{}
    decoder := xml.NewDecoder(bytes.NewReader(data))
    encoder := xml.NewEncoder(b)
    encoder.Indent("", "  ")
    for {
        token, err := decoder.Token()
        if err == io.EOF {
            encoder.Flush()
            return b.Bytes(), nil
        }
        if err != nil {
            return nil, err
        }
        err = encoder.EncodeToken(token)
        if err != nil {
            return nil, err
        }
    }
}