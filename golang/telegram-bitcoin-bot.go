// ==================== ./main.go
package main

import (
	"log"
	"time"

	"github.com/tomassirio/bitcoinTelegram/config"
	"github.com/tomassirio/bitcoinTelegram/handler"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	bot, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		Token:  config.LoadConfig().Token, // os.Getenv("TOKEN")
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	for k, v := range handler.LoadHandler(bot) {
		bot.Handle(k, v)
		log.Println(k + "✅ Loaded!")
	}

	bot.Start()
}

// ==================== ./config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
}

func LoadConfig() *Config {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{Token: os.Getenv("TOKEN")}
}

// ==================== ./handler/handler.go
package handler

import (
	"github.com/tomassirio/bitcoinTelegram/commands"
	tb "gopkg.in/tucnak/telebot.v2"
)

func LoadHandler(b *tb.Bot) map[string]func(m *tb.Message) {
	commandMap := make(map[string]func(m *tb.Message))

	commandMap["/price"] = func(m *tb.Message) {
		res, _ := commands.GetPrice()
		b.Send(m.Chat, "BTC's Current price is: U$S "+res)
	}

	commandMap["/historic"] = func(m *tb.Message) {
		res, g, _ := commands.GetHistoric()
		b.Send(m.Chat, "BTC's Price compared to yesterday is: "+res)
		b.Send(m.Chat, g)
	}

	commandMap["/summary"] = func(m *tb.Message) {
		p, h, _ := commands.GetSummary()
		b.Send(m.Chat, "BTC's Current price is: U$S "+p+"\nBTC's Price compared to yesterday is: "+h)
	}

	return commandMap
}

// ==================== ./utils/apiCall.go
package utils

import (
	"encoding/json"
	"net/http"

	"github.com/tomassirio/bitcoinTelegram/model"
)

func GetApiCall() (*model.Price, error) {
	resp, err := http.Get("https://bitex.la/api-v1/rest/btc_usd/market/ticker")
	p := &model.Price{}

	if err != nil {
		return p, err
	}

	err = json.NewDecoder(resp.Body).Decode(p)
	return p, err
}

// ==================== ./commands/historic.go
package commands

import (
	"fmt"
	"math"

	tb "gopkg.in/tucnak/telebot.v2"

	"github.com/tomassirio/bitcoinTelegram/utils"
)

func GetHistoric() (string, *tb.Animation, error) {
	p, err := utils.GetApiCall()
	l := p.Last
	o := p.Open
	his := ((l - o) / o) * 100
	if !math.Signbit(float64(his)) {
		g := &tb.Animation{File: tb.FromURL(
               "https://i.pinimg.com/originals/e4/38/99/e4389936b099672128c54d25c4560695.gif")}
		return "%" + fmt.Sprintf("%.2f", ((l-o)/o)*100), g, err
	} else {
		g := &tb.Animation{File: tb.FromURL(
               "http://www.brainlesstales.com/bitcoin-assets/images/fan-versions/2015-01-osEroUI.gif")}
		return "-%" + fmt.Sprintf("%.2f", -1*((l-o)/o)*100), g, err
	}
}

// ==================== ./commands/price.go
package commands

import (
	"fmt"

	"github.com/tomassirio/bitcoinTelegram/utils"
)

func GetPrice() (string, error) {
	p, err := utils.GetApiCall()
	return fmt.Sprintf("%.2f", p.Last), err
}

// ==================== ./commands/summary.go
package commands

import (
	"fmt"
	"math"

	"github.com/tomassirio/bitcoinTelegram/utils"
)

func GetSummary() (string, string, error) {
	p, err := utils.GetApiCall()
	l := p.Last
	o := p.Open
	his := ((l - o) / o) * 100
	if !math.Signbit(float64(his)) {
		return fmt.Sprintf("%.2f", p.Last), "%" + fmt.Sprintf("%.2f", ((l-o)/o)*100), err
	} else {
		return fmt.Sprintf("%.2f", p.Last), "-%" + fmt.Sprintf("%.2f", -1*((l-o)/o)*100), err
	}
}

// ==================== ./model/price.go
package model

type Price struct {
	Last            float32 `json:"last"`
	PriceBeforeLast float32 `json:"price_before_last"`
	Open            float32 `json:"open"`
	High            float32 `json:"high"`
	Low             float32 `json:"low"`
	Vwap            float32 `json:"vwap"`
	Volume          float32 `json:"volume"`
	Bid             float32 `json:"bid"`
	Ask             float32 `json:"ask"`
}

// https://bitex.la/developers#api_ticker
/*
{
  "last":               1230.0,  // Last transaction price
  "price_before_last":  1220.0,  // Helps you tell if price is going up or down.
  "open":        1198.45875559,  // What the price was 24 hours ago.
  "high":               1230.0,  // Highest price for the past 24 hours.
  "low":          1193.2507548,  // Lowest price for the past 24 hours.
  "vwap":        1208.57944642,  // Volume-Weighted Average Price for the past 24 hours.
  "volume":        16.45315074,  // Transacted volume for the last 24 hours.
  "bid":          1226.5583985,  // Highest current buy order.
  "ask":         1235.71481927   // Lowest current ask order.
}
*/
