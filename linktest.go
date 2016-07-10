package main

import (
  "fmt"
  "strings"
  "os"
  "github.com/urfave/cli"
  "github.com/PuerkitoBio/goquery"
)

type BrokenLink struct {
  source string
  destination string
}

func main() {

  foundUrls := make(map[string]bool)
  brokenUrls := make(map[BrokenLink]bool)
  crawlsStarted := 0
  crawlsFinished := 0

  // Channels
  chUrls := make(chan string)
  chBrokenUrls := make(chan BrokenLink)
  chCrawlsStarted := make(chan bool)
  chCrawlsFinished := make(chan bool)

  // Initialize the CLI appliaction
  app := cli.NewApp()
  app.Name = "linktest"
  app.Usage = "Crawl a website, searching for broken links and imgs."
  app.Version = "v0.0.1"
  app.Action = func(c *cli.Context) error {
    crawlsStarted++
    seedUrl := c.Args().First()
    go crawl("seed", seedUrl, chUrls, chBrokenUrls, chCrawlsStarted, chCrawlsFinished)

    // Subscribe to all channels
    for crawlsStarted > crawlsFinished {
      select {
        case url := <- chUrls:
          foundUrls[url] = true
        case brokenUrl := <- chBrokenUrls:
          brokenUrls[brokenUrl] = true
        case <- chCrawlsStarted:
          crawlsStarted++
        case <- chCrawlsFinished:
          crawlsFinished++
      }
    }

    fmt.Println("Links found:")

    for url, _ := range foundUrls {
        fmt.Println(url)
    }

    fmt.Println("\nBroken links found:")

    for url, _ := range brokenUrls {
        fmt.Println("Source: " + url.source + "  Destination: " + url.destination)
    }

    return nil
  }

  app.Run(os.Args)
}

func crawl(sourceUrl string, destinationUrl string, chUrls chan string, chBrokenUrls chan BrokenLink, chStarted chan bool, chFinished chan bool) {

  fmt.Println("Crawling " + destinationUrl)

  if sourceUrl != "seed" {
    chStarted <- true
  }

  defer func() {
    // Notify that we're done after this function
    chFinished <- true
  }()

  doc, err := goquery.NewDocument(destinationUrl)

  if err != nil {
    b := BrokenLink{sourceUrl, destinationUrl}
    chBrokenUrls <- b
  }

  // Don't crawl if it starts with an http
  hasProto := strings.HasPrefix(destinationUrl, "http")
  if hasProto && sourceUrl != "seed" {
    return
  }

  doc.Find("a").Each(func(i int, s *goquery.Selection) {
    link, exists := s.Attr("href")
    if exists {
      chUrls <- link
    }
  })
}
