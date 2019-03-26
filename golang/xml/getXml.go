// tweaked from: https://stackoverflow.com/a/42718113/1170664
func getXML(url string) ([]byte, error) {
  resp, err := http.Get(url)
  if err != nil {
    return []byte{}, fmt.Errorf("GET error: %v", err)
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return []byte{}, fmt.Errorf("Read body: %v", err)
  }

  return data, nil
}

if xmlBytes, err := getXML("http://somehost.com/some.xml"); err != nil {
  log.Printf("Failed to get XML: %v", err)
} else {
  var result myXMLstruct
  xml.Unmarshal(xmlBytes, &result)
}