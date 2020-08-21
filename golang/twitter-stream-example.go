// https://github.com/Fallenstedt/twitter-stream/tree/master/example

// main.go
package main

const KEY = "KEY"
const SECRET = "SECRET"

func main() {
	// Run an example function
	addRules()
	getRules()
	initiateStream()
	//deleteRules()
}

// create_rules_example.go
package main

import (
	"fmt"
	twitterstream "github.com/fallenstedt/twitter-stream"
	"github.com/fallenstedt/twitter-stream/rules"
)

func addRules() {

	tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret(KEY, SECRET).RequestBearerToken()
	if err != nil {
		panic(err)
	}
	api := twitterstream.NewTwitterStream(tok.AccessToken)
	rules := twitterstream.NewRuleBuilder().
		AddRule("cat has:images", "cat tweets with images").
		AddRule("puppy has:images", "puppy tweets with images").
		AddRule("lang:en -is:retweet -is:quote (#golangjobs OR #gojobs)", "golang jobs").
		Build()

	res, err := api.Rules.Create(rules, false) // dryRun is set to false.
	if err != nil {
		panic(err)
	}

	if res.Errors != nil && len(res.Errors) > 0 {
		//https://developer.twitter.com/en/support/twitter-api/error-troubleshooting
		panic(fmt.Sprintf("Received an error from twitter: %v", res.Errors))
	}

	fmt.Println("I have created rules.")
	printRules(res.Data)
}

func getRules() {
	tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret(KEY, SECRET).RequestBearerToken()
	if err != nil {
		panic(err)
	}
	api := twitterstream.NewTwitterStream(tok.AccessToken)
	res, err := api.Rules.Get()
	if err != nil {
		panic(err)
	}

	if res.Errors != nil && len(res.Errors) > 0 {
		//https://developer.twitter.com/en/support/twitter-api/error-troubleshooting
		panic(fmt.Sprintf("Received an error from twitter: %v", res.Errors))
	}

	if len(res.Data) > 0 {
		fmt.Println("I found these rules: ")
		printRules(res.Data)
	} else {
		fmt.Println("I found no rules")
	}
}

func deleteRules() {
	tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret(KEY, SECRET).RequestBearerToken()
	if err != nil {
		panic(err)
	}
	api := twitterstream.NewTwitterStream(tok.AccessToken)

	// use api.Rules.Get to find the ID number for an existing rule
	res, err := api.Rules.Delete(rules.NewDeleteRulesRequest(1469777072675450881, 74893274932), false)
	if err != nil {
		panic(err)
	}

	if res.Errors != nil && len(res.Errors) > 0 {
		//https://developer.twitter.com/en/support/twitter-api/error-troubleshooting
		panic(fmt.Sprintf("Received an error from twitter: %v", res.Errors))
	}

	fmt.Println("I have deleted rules ")
}

func printRules(data []rules.DataRule) {
	for _, datum := range data {
		fmt.Printf("Id: %v\n", datum.Id)
		fmt.Printf("Tag: %v\n",datum.Tag)
		fmt.Printf("Value: %v\n\n", datum.Value)
	}
}

// stream_forever.go
package main

import (
	"encoding/json"
	"fmt"
	twitterstream "github.com/fallenstedt/twitter-stream"
	"github.com/fallenstedt/twitter-stream/stream"
	"time"
)

// This example assumes you have atleast 1 twitter rule created.
// See "create_rules_example.go" to create a rule.

// Establishing a connection to the streaming APIs means making a very long lived HTTPS request, and parsing the response incrementally.
// When connecting to the sampled stream endpoint, you should form a HTTPS request and consume the resulting stream for as long as is practical.
// Twitter servers will hold the connection open indefinitely, barring server-side error, excessive client-side lag, network issues, routine server maintenance, or duplicate logins.
// With connections to streaming endpoints, **it is likely, and should be expected,** that disconnections will take place and reconnection logic built.
// ~https://developer.twitter.com/en/docs/twitter-api/tweets/volume-streams/integrate/handling-disconnections

type StreamDataExample struct {
	Data struct {
		Text      string    `json:"text"`
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		AuthorID  string    `json:"author_id"`
	} `json:"data"`
	Includes struct {
		Users []struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"users"`
	} `json:"includes"`
	MatchingRules []struct {
		ID  string `json:"id"`
		Tag string `json:"tag"`
	} `json:"matching_rules"`
}

// This will run forever
func initiateStream() {
	fmt.Println("Starting Stream")

	// Start the stream
	// And return the library's api
	api := fetchTweets()

	// When the loop below ends, restart the stream
	defer initiateStream()

	// Start processing data from twitter
	for tweet := range api.GetMessages() {

		// Handle disconnections from twitter
		// https://developer.twitter.com/en/docs/twitter-api/tweets/volume-streams/
		// integrate/handling-disconnections
		if tweet.Err != nil {
			fmt.Printf("got error from twitter: %v", tweet.Err)

			// Notice we "StopStream" and then "continue" the loop instead of breaking.
			// StopStream will close the long running GET request to Twitter's v2 Streaming
			// endpoint by closing the `GetMessages` channel. Once it's closed, it's safe to
			// perform a new network request with `StartStream`
			api.StopStream()
			continue
		}
		result := tweet.Data.(StreamDataExample)

		// Here I am printing out the text.
		// You can send this off to a queue for processing.
		// Or do your processing here in the loop
		fmt.Println(result.Data.Text)
	}

	fmt.Println("Stopped Stream")
}

func fetchTweets() stream.IStream {
	// Get Bearer Token using API keys
	tok, err := getTwitterToken()
	if err != nil {
		panic(err)
	}

	// Instantiate an instance of twitter stream using the bearer token
	api := getTwitterStreamApi(tok)

	// On Each tweet, decode the bytes into a StreamDataExample struct
	api.SetUnmarshalHook(func(bytes []byte) (interface{}, error) {
		data := StreamDataExample{}
		if err := json.Unmarshal(bytes, &data); err != nil {
			fmt.Printf("failed to unmarshal bytes: %v", err)
		}
		return data, err
	})

	// Request additional data from teach tweet
	streamExpansions := twitterstream.NewStreamQueryParamsBuilder().
		AddExpansion("author_id").
		AddTweetField("created_at").
		Build()

	// Start the Stream
	err = api.StartStream(streamExpansions)
	if err != nil {
		panic(err)
	}

	// Return the twitter stream api instance
	return api
}

func getTwitterToken() (string, error) {
	tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret(KEY, SECRET).RequestBearerToken()
	return tok.AccessToken, err
}

func getTwitterStreamApi(tok string) stream.IStream {
	return twitterstream.NewTwitterStream(tok).Stream
}