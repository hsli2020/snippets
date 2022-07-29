A Very Simple Tech Stack: Building a small app with HTMX, Go Fiber and Uno CSS
https://github.com/awesome-club/htmx-go

// stock.go
package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2/log"
)

const PoligonPath = "https://api.polygon.io"
const ApiKey = "apiKey=<your-key>"

const TickerPath = PoligonPath + "/v3/reference/tickers"
const DailyValuesPath = PoligonPath + "/v1/open-close"

type Stock struct {
	Ticker string `json:"ticker"`
	Name   string `json:"name"`
}

type SearchResult struct {
	Results []Stock `json:"results"`
}

type Values struct {
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
}

func Fetch(path string) []byte {
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func SearchTicker(ticker string) []Stock {
	body := Fetch(TickerPath + "?" + ApiKey + "&ticker=" + strings.ToUpper(ticker))

	data := SearchResult{}
	json.Unmarshal(body, &data)

	return data.Results
}

func GetDailyValues(ticker string) Values {
	body := Fetch(DailyValuesPath + "/" + strings.ToUpper(ticker) + "/2023-09-15/?" + ApiKey)

	data := Values{}
	json.Unmarshal(body), &data)

	return data
}

// server.go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/search", func(c *fiber.Ctx) error {
		ticker := c.Query("ticker")
		results := SearchTicker(ticker)

		return c.Render("results", fiber.Map{
			"Results": results,
		})
	})

	app.Get("/values/:ticker", func(c *fiber.Ctx) error {
		ticker := c.Params("ticker")
		values := GetDailyValues(ticker)

		return c.Render("values", fiber.Map{
			"Ticker": ticker,
			"Values": values,
		})
	})

	app.Listen(":3000")
}

// view/index.html
<html>
  <head>
    <script src="https://unpkg.com/htmx.org@1.9.4"></script>
    <script src="https://cdn.jsdelivr.net/npm/@unocss/runtime"></script>
  </head>
  <body>
    <main class="container mx-auto w-sm">
      <input class="rounded-lg mx-auto w-full p-4 shadow-lg font-size-5 font-semibold"
        type="text"
        name="ticker"
        hx-get="/search"
        hx-trigger="keyup changed delay:100ms"
        hx-target="#search-results"
      />
      <div id="search-results"></div>

      <div id="stock-details"></div>
    </main>
  </body>
</html>

// view/results.html
<ul class="list-none m0 mt-2 p0 rounded-lg shadow-xl border">
  {{range .Results}}
    <li>
      <button
        class="bg-white b-none block p3 w-full rounded-lg font-size-4 text-left hover:bg-#fafafa"
        hx-get="/values/{{.Ticker}}"
        hx-trigger="click"
        hx-target="#stock-details">
        {{.Name}}
      </button>
    </li>
  {{end}}
</ul>

// view/values.html
<h3>{{.Ticker}}</h3>
<ul class="flex border-#eee m0 p0">
  <li class="flex flex-col flex-1"><strong>High</strong>{{.Values.High}}</li>
  <li class="flex flex-col flex-1"><strong>Open</strong>${{.Values.Open}}</li>
  <li class="flex flex-col flex-1"><strong>Low</strong>${{.Values.Low}}</li>
</ul>
