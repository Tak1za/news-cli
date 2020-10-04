package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tak1za/go-news/config"
	"github.com/Tak1za/go-news/pkg/news"
	"github.com/urfave/cli"
)

func main() {
	config := config.New()
	app := &cli.App{
		Name:  "news",
		Usage: "get latest news",
		Action: func(c *cli.Context) error {
			fmt.Println("Getting news has never been so easy!!")
			return nil
		},
		Commands: []cli.Command{
			{
				Name:    "top",
				Aliases: []string{"t"},
				Usage:   "Get top news",
				Action: func(c *cli.Context) error {
					articles := news.GetTopNews(config)
					for _, v := range articles {
						fmt.Println("Title: ", v.Title)
						fmt.Println("Content: ", v.Content)
						fmt.Println("Full Story: ", v.Url)
						fmt.Println()
						fmt.Println("---------xxx---------")
						fmt.Println()
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
