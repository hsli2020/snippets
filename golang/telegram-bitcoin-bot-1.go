package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	bot, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range GetHandlers(bot) {
		bot.Handle(k, v)
		log.Println(k + "✅ Loaded!")
	}

	bot.Start()
}

func GetHandlers(b *tb.Bot) map[string]func(m *tb.Message) {
	commandMap := make(map[string]func(m *tb.Message))

	commandMap["/price"] = func(m *tb.Message) {
		res, _ := GetPrice()
		b.Send(m.Chat, "BTC's Current price is: US$ "+res)
	}

	commandMap["/historic"] = func(m *tb.Message) {
		res, g, _ := GetHistoric()
		b.Send(m.Chat, "BTC's Price compared to yesterday is: "+res)
		b.Send(m.Chat, g)
	}

	commandMap["/summary"] = func(m *tb.Message) {
		p, h, _ := GetSummary()
		b.Send(m.Chat, "BTC's Current price is: U$S "+p+"\nBTC's Price compared to yesterday is: "+h)
	}

	return commandMap
}

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

func GetApiCall() (*Price, error) {
	resp, err := http.Get("https://bitex.la/api-v1/rest/btc_usd/market/ticker")
	p := &Price{}

	if err != nil {
		return p, err
	}

	err = json.NewDecoder(resp.Body).Decode(p)
	return p, err
}

func GetHistoric() (string, *tb.Animation, error) {
	p, err := GetApiCall()
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

func GetPrice() (string, error) {
	p, err := GetApiCall()
	return fmt.Sprintf("%.2f", p.Last), err
}

func GetSummary() (string, string, error) {
	p, err := GetApiCall()
	l := p.Last
	o := p.Open
	his := ((l - o) / o) * 100
	if !math.Signbit(float64(his)) {
		return fmt.Sprintf("%.2f", p.Last), "%" + fmt.Sprintf("%.2f", ((l-o)/o)*100), err
	} else {
		return fmt.Sprintf("%.2f", p.Last), "-%" + fmt.Sprintf("%.2f", -1*((l-o)/o)*100), err
	}
}
