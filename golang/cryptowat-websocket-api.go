package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// Initialize a connection using your API key
// You can generate an API key here: https://cryptowat.ch/account/api-access
// Paste your API key here:
const (
	APIKEY = "NZAP2FY3HMEC5T1BOYA3"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("wss://stream.cryptowat.ch/connect?apikey="+APIKEY, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Read first message, which should be an authentication response
	_, message, err := c.ReadMessage()
	var authResult struct {
		AuthenticationResult struct {
			Status string `json:"status"`
		} `json:"authenticationResult"`
	}
	err = json.Unmarshal(message, &authResult)
	if err != nil {
		panic(err)
	}

	// Send a JSON payload to subscribe to a list of resources
	// Read more about resources here:
	// https://docs.cryptowat.ch/websocket-api/data-subscriptions#resources
	resources := []string{
		"instruments:9:trades",
	}
	subMessage := struct {
		Subscribe SubscribeRequest `json:"subscribe"`
	}{}

	// No map function in golang :-(
	for _, resource := range resources {
		subMessage.Subscribe.Subscriptions = append(subMessage.Subscribe.Subscriptions,
			Subscription{StreamSubscription: StreamSubscription{Resource: resource}})
	}
	msg, err := json.Marshal(subMessage)
	err = c.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		panic(err)
	}

	// Process incoming BTC/USD trades
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal("Error reading from connection", err)
			return
		}

		var update Update
		err = json.Unmarshal(message, &update)
		if err != nil {
			panic(err)
		}

		for _, trade := range update.MarketUpdate.TradesUpdate.Trades {
			log.Printf(
				"BTC/USD trade on market %d: %s %s",
				update.MarketUpdate.Market.MarketId,
				trade.Price,
				trade.Amount,
			)
		}
	}
}

// Helper types for JSON serialization

type Subscription struct {
	StreamSubscription `json:"streamSubscription"`
}

type StreamSubscription struct {
	Resource string `json:"resource"`
}

type SubscribeRequest struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

type Update struct {
	MarketUpdate struct {
		Market struct {
			MarketId int `json:"marketId,string"`
		} `json:"market"`

		TradesUpdate struct {
			Trades []Trade `json:"trades"`
		} `json:"tradesUpdate"`
	} `json:"marketUpdate"`
}

type Trade struct {
	Timestamp     int `json:"timestamp,string"`
	TimestampNano int `json:"timestampNano,string"`

	Price  string `json:"priceStr"`
	Amount string `json:"amountStr"`
}
