package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
)

func HashSignature(str string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(str))
	if err != nil {
		return ""
	}

	hash := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	hash = url.QueryEscape(hash)
	return hash
}

func main() {
	/* here is the canonical string from Amazon documentation which should yield the expected signature below
	GET
	webservices.amazon.com
	/onca/xml
	AWSAccessKeyId=AKIAIOSFODNN7EXAMPLE&AssociateTag=mytag-20&ItemId=0679722769&Operation=ItemLookup&ResponseGroup=Images%2CItemAttributes%2COffers%2CReviews&Service=AWSECommerceService&Timestamp=2014-08-18T12%3A00%3A00Z&Version=2013-08-01
	*/

	SECRET_KEY := "1234567890"
	CANONICAL_STR := "GET\nwebservices.amazon.com\n/onca/xml\nAWSAccessKeyId=AKIAIOSFODNN7EXAMPLE&AssociateTag=mytag-20&ItemId=0679722769&Operation=ItemLookup&ResponseGroup=Images%2CItemAttributes%2COffers%2CReviews&Service=AWSECommerceService&Timestamp=2014-08-18T12%3A00%3A00Z&Version=2013-08-01"

	EXPECTED := "j7bZM0LXZ9eXeZruTqWm2DIvDYVUU3wxPPpp%2BiXxzQc%3D"

	if RESULT := HashSignature(CANONICAL_STR, SECRET_KEY); RESULT != EXPECTED {
		fmt.Errorf("\nEXPECTED:\n%v\nRESULT:\n%v", EXPECTED, RESULT)
	} else {
		fmt.Println("TestAmazonSignature: Signature: OK")
		fmt.Println(HashSignature(CANONICAL_STR, SECRET_KEY))
	}
}
