package main

import "fmt"
import "crypto/hmac"
import "crypto/sha256"
import "time"
import "strings"
import "encoding/base64"
import "net/http"
import "net/url"
import "io/ioutil"

const METHOD ="GET"
const HOST ="webservices.amazon.com"
const URI= "/onca/xml"
const QUERY_STRING= "AWSAccessKeyId=121212121212&AssociateTag=smasholab-20&IdType=ISBN&ItemId=B000MQTJW2&Operation=ItemLookup&Service=AWSECommerceService&Timestamp=%s"


func main() {
	t := time.Now()
	tm:= t.Format("2006-01-02T15:04:05Z")
	tm= url.QueryEscape(tm)
	fmt.Println("tm:", tm)

	query := fmt.Sprintf(QUERY_STRING, tm)
	ul := fmt.Sprintf(QUERY_STRING, tm)
	fmt.Println("query:", query)


	//AWSAccessKeyId := "sssbbbsssbbb"
	AWSSecretKeyId := "ooxxooxx"

	sha256 := sha256.New
	hash := hmac.New(sha256, []byte(AWSSecretKeyId))
	template:= "%s\n%s\n%s\n%s"
	template= fmt.Sprintf(template, METHOD, HOST, URI, query)
	fmt.Println("template:", template)
	hash.Write([]byte(template))
	sha := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	sha= url.QueryEscape(sha)
	fmt.Println("sha", sha)			
	
	ul=  ul + "&Signature=" +sha
	ul= "http://webservices.amazon.com/onca/xml?"+ ul 

	ul= strings.Replace(ul, "+", "%20", -1)
	ul= strings.Replace(ul, "*", "%2A", -1)
	ul= strings.Replace(ul, "%7E", "~", -1)


	fmt.Println("url:", ul)


	//request
	response, err := http.Get(ul)
	if err != nil {
		fmt.Println("err", err)
	}else{
		content, _ := ioutil.ReadAll(response.Body) 
		println("response", string(content))
		response.Body.Close()
	}	
}
