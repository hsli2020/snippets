package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
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

var sample1 = `<?xml version="1.0" encoding="UTF-8" ?>
<XMLFORMPOST><REQUEST>price-availability</REQUEST><LOGIN><USERID>800712USR</USERID><PASSWORD>USR@XML2013</PASSWORD></LOGIN><PARTNUM>980001203CDN</PARTNUM><PARTNUM>XS708T100NESCA</PARTNUM></XMLFORMPOST>`

var sample2 = `<?xml version="1.0" encoding="UTF-8" ?>
<XMLRESPONSE><ITEM><PARTNUM>980001203CDN</PARTNUM><UNITPRICE>88.39</UNITPRICE><BRANCHQTY><BRANCH>Vancouver</BRANCH><QTY>6</QTY><INSTOCKDATE></INSTOCKDATE></BRANCHQTY><BRANCHQTY><BRANCH>Toronto</BRANCH><QTY>0</QTY><INSTOCKDATE></INSTOCKDATE></BRANCHQTY><TOTALQTY>6</TOTALQTY></ITEM><ITEM><PARTNUM>XS708T100NESCA</PARTNUM><UNITPRICE>843.21</UNITPRICE><BRANCHQTY><BRANCH>Vancouver</BRANCH><QTY>0</QTY><INSTOCKDATE></INSTOCKDATE></BRANCHQTY><BRANCHQTY><BRANCH>Toronto</BRANCH><QTY>1</QTY><INSTOCKDATE></INSTOCKDATE></BRANCHQTY><TOTALQTY>1</TOTALQTY></ITEM><STATUS>success</STATUS></XMLRESPONSE>
`

func main() {
	out1, _ := formatXML([]byte(sample1))
	fmt.Println(string(out1))

	out2, _ := formatXML([]byte(sample2))
	fmt.Println(string(out2))
}
