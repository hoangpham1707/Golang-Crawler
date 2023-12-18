package main

import (
	web "crawler/web"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {

	web.StartCrawler()
	c := cron.New()
	c.AddFunc("@every 10h", func() {
		fmt.Println("Starting data crawl...")
		web.StartCrawler()
		fmt.Println("Data crawl completed.")
	})
	c.Start()

	time.Sleep(30 * time.Minute)

}
