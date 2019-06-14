/*
This code Do the following stuff :
1) login to a website called website.com by submitting password and username on the page with url:
   - http://website.com/login
2) Now after login using the cookies stored by this webiste access user profile page
3) Now using same client which stored the required cookies make another post request to user
   profile page present at page :- http://website.com/upser_profile_page .
4) Now get html of this whole page and print it in log as a string .
*/
package main

import (
	"code.google.com/p/go.net/publicsuffix"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func main() {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{Jar: jar}
	resp, err := client.PostForm("http://website.com/login", url.Values{
		"password": {"loginpassword"},
		"username": {"testuser"},
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err = client.PostForm("http://website.com/upser_profile_page", url.Values{
		"userid": {"2"},
	})
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data)) // print whole html of user profile data
}
