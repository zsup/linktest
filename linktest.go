package main

import (
  "fmt"
  "log"
  // "io/ioutil"
  // "net/http"
  "os"
  "github.com/urfave/cli"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  app := cli.NewApp()
  app.Name = "linktest"
  app.Usage = "Crawl a website, searching for broken links and imgs."
  app.Version = "v0.0.1"
  app.Action = func(c *cli.Context) error {
    url := c.Args().First()
    scrape(url)
    return nil
  }

  app.Run(os.Args)
}

func scrape(url string) {
  doc, err := goquery.NewDocument(url)

  if err != nil {
    log.Fatal(err)
  }

  doc.Find("a").Each(func(i int, s *goquery.Selection) {
    link, exists := s.Attr("href")
    if exists {
      fmt.Println(link)
    }
  })
}
