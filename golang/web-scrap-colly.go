package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	stocks := []string{	"AAPL",	"MSFT",	"PLTR",	"AMZN"	}

	parseStocks(stocks)
	fmt.Println("Completed the code process, took: %f seconds", time.Since(start).Seconds())
}

func parseStocks(stocks []string) {
	ch := make(chan string)
	var wg sync.WaitGroup

	for _, stock := range stocks {
		wg.Add(1)
		go parseStock(stock, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}

func parseStock(stock string, ch chan string, wg *sync.WaitGroup) {
	defer (*wg).Done()

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"),
		colly.AllowedDomains("finance.yahoo.com"),
		colly.MaxBodySize(0),
		colly.AllowURLRevisit(),
		colly.Async(true),
	)

	// Set max Parallelism and introduce a Random Delay
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			dataSlice := []string{}
			el.ForEach("td", func(_ int, el *colly.HTMLElement) {
				dataSlice = append(dataSlice, el.Text)
			})

			if dataSlice[0] == "Previous Close" {
				ch <- stock + " Price for previous close is: " + dataSlice[1]
			}
		})
	})

	c.Visit("https://finance.yahoo.com/quote/" + stock)
	c.Wait()
}
