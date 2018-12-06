package gomws

import "net/url"
import "net/http"
import "time"
import "strings"
import "crypto/sha256"
import "crypto/hmac"
import "encoding/base64"
import "fmt"
import "os"
import "errors"

func init() {
	UserAgent = (func() string {
		h, _ := os.Hostname()

		return fmt.Sprintf("%s/gomws (Language=go; Host=%s)", CompanyName, h)
	})()
}

var UserAgent string
var CompanyName = "github.com/shanemhansen"
var IncompleteRequest = errors.New("incomplete request")

type Creds struct {
	Country   string
	AccessId  string
	AccessKey string
	Merchant  string
}
type Client struct {
	Creds
	Method           string // GET, PUT, POST, etc.
	Region           Region
	Action           string
	Parameters       url.Values
	SignatureMethod  string
	SignatureVersion string
	CompanyName      string
}

func NewClient(client Client) Client {
	client.SignatureVersion = "2"
	client.SignatureMethod = "HmacSHA256"
	client.Region = RegionByCountry(client.Country)
	if client.Parameters == nil {
		client.Parameters = make(url.Values)
	}
	return client
}

func (this *Client) Request() (req *http.Request, err error) {
	if this.Method == "" || this.AccessId == "" || this.AccessKey == "" || this.Merchant == "" ||
		this.Region.Endpoint == "" {
		err = IncompleteRequest
		return
	}
	url, err := url.Parse(this.Region.Endpoint)
	if err != nil {
		return
	}
	this.Parameters.Add("Merchant", this.Merchant)
	this.Parameters.Add("AWSAccessKeyId", this.AccessId)
	this.Parameters.Add("SignatureMethod", this.SignatureMethod)
	this.Parameters.Add("SignatureVersion", this.SignatureVersion)
	this.Parameters.Add("Version", "2009-01-01")
	this.Parameters.Add("Action", this.Action)
	this.Parameters.Add("Timestamp", XMLTimestamp(time.Now()))
	stringToSign, err := this.StringToSign()
	if err != nil {
		return
	}
	signature := Sign(stringToSign, []byte(this.AccessKey))
	this.Parameters.Add("Signature", signature)
	url.RawQuery = CanonicalizedQueryString(this.Parameters)
	req, err = http.NewRequest(this.Method, url.String(), nil)
	req.Header.Add("User-Agent", UserAgent)
	return
}

var ISO8601 = "2006-01-02T15:04:05Z"

func XMLTimestamp(t time.Time) string {
	return t.UTC().Format(ISO8601)
}

func CanonicalizedQueryString(values url.Values) (str string) {
	// per aws docs and docs for values.Encode, we respect RFC 3986
	// we may not deal with utf-8, only ascii
	// params are sorted
	// we have to fix the '+' to '%20'
	str = values.Encode()
	str = strings.Replace(str, "+", "%20", -1)
	return
}

func (this *Client) StringToSign() (stringToSign string, err error) {
	endpoint, err := url.Parse(this.Region.Endpoint)
	if err != nil {
		return
	}
	stringToSign = strings.Join([]string{
		this.Method,
		strings.ToLower(endpoint.Host),
		endpoint.Path,
		CanonicalizedQueryString(this.Parameters),
	}, "\n")

	return
}

func Sign(str string, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
