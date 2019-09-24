package main

// https://play.golang.org/p/JUqQY3WpW5

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strings"
)

const flatxml = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://example.com/ns"><soapenv:Header/><soapenv:Body><ns:request><ns:customer><ns:id>123</ns:id><ns:name type="NCHZ">John Brown</ns:name></ns:customer></ns:request></soapenv:Body></soapenv:Envelope>`

func main() {
	buf := new(bytes.Buffer)
	d := xml.NewDecoder(strings.NewReader(flatxml))
	e := xml.NewEncoder(buf)
	e.Indent("", " ")

tokenize:
	for {
		tok, err := d.Token()
		switch {
		case err == io.EOF:
			e.Flush()
			break tokenize
		case err != nil:
			log.Fatal(err)
		}
		e.EncodeToken(tok)
	}

	newxml := buf.String()
	fmt.Println(newxml)
}
