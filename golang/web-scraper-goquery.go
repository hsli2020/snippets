package main

import (
    // import standard libraries
    "fmt"
    "log"

    // import third party libraries
    "github.com/PuerkitoBio/goquery"
)

func postScrape() {
    doc, err := goquery.NewDocument("http://jonathanmh.com")
    if err != nil {
        log.Fatal(err)
    }

    // use CSS selector found with the browser inspector
    // for each, use index and item
    doc.Find("#main article .entry-title").Each(func(index int, item *goquery.Selection) {
        title := item.Text()
        linkTag := item.Find("a")
        link, _ := linkTag.Attr("href")
        fmt.Printf("Post #%d: %s - %s\n", index, title, link)
    })
}

func linkScrape() {
    doc, err := goquery.NewDocument("http://jonathanmh.com")
    if err != nil {
        log.Fatal(err)
    }

    // use CSS selector found with the browser inspector
    // for each, use index and item
    doc.Find("body a").Each(func(index int, item *goquery.Selection) {
        linkTag := item
        link, _ := linkTag.Attr("href")
        linkText := linkTag.Text()
        fmt.Printf("Link #%d: '%s' - '%s'\n", index, linkText, link)
    })
}

func metaScrape() {
    doc, err := goquery.NewDocument("http://jonathanmh.com")
    if err != nil {
        log.Fatal(err)
    }

    var metaDescription string
    var pageTitle string

    // use CSS selector found with the browser inspector
    // for each, use index and item
    pageTitle = doc.Find("title").Contents().Text()

    doc.Find("meta").Each(func(index int, item *goquery.Selection) {
        if( item.AttrOr("name","") == "description") {
            metaDescription = item.AttrOr("content", "")
        }
    })
    fmt.Printf("Page Title: '%s'\n", pageTitle)
    fmt.Printf("Meta Description: '%s'\n", metaDescription)
}

func main() {
    postScrape()
    linkScrape()
    metaScrape()
}