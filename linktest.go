package main

import (
  "fmt"
  "os"
  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "linktest"
  app.Usage = "Crawl a website, searching for broken links and imgs."
  app.Action = func(c *cli.Context) error {
    fmt.Printf("Hello %q", c.Args().First())
    return nil
  }

  app.Run(os.Args)
}
