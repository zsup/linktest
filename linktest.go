package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "linktest"
  app.Usage = "Crawl a website, searching for broken links and imgs."
  app.Version = "v0.0.1"
  app.Action = func(c *cli.Context) error {
    url := c.Args().First()
    resp, err := http.Get(url)

    if err != nil {
        fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
        return nil
    }

    bytes, _ := ioutil.ReadAll(resp.Body)

    fmt.Println("HTML:\n\n", string(bytes))

    resp.Body.Close()

    return nil
  }

  app.Run(os.Args)
}
