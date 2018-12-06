package main

import "fmt"
import "encoding/xml"
import "encoding/json"

type User struct {
    XMLName xml.Name `xml:"users"`

    List []struct {
        //XMLName xml.Name `xml:"user"`         THIS DOESN'T WORK!
        Type string `xml:"type,attr"`
        Name string `xml:"name"`

        Social struct {
            //XMLName xml.Name `xml:"social"`   THIS DOESN'T WORK!
            Facebook string `xml:"facebook"`
            Twitter string `xml:"twitter"`
            Youtube string `xml:"youtube"`
        } `xml:"social"`                        THIS WORKS
    } `xml:"user"`                              THIS WORKS
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func main() {
	data := `<?xml version="1.0" encoding="UTF-8"?>
<users>
  <user type="admin">
    <name>Elliot</name>
    <social>
      <facebook>https://facebook.com</facebook>
      <twitter>https://twitter.com</twitter>
      <youtube>https://youtube.com</youtube>
    </social>
  </user>  
  <user type="reader">
    <name>Fraser</name>
    <social>
      <facebook>https://facebook.com</facebook>
      <twitter>https://twitter.com</twitter>
      <youtube>https://youtube.com</youtube>
    </social>
  </user>  
</users>
	`
	var user User;

	err := xml.Unmarshal([]byte(data), &user)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	PrettyPrint(user)
}
