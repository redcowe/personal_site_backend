package spider

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type Book struct {
	title  string
	author string
	link   string
}

func Crawl(target string) {
	c := colly.NewCollector(colly.AllowedDomains(
		"https://bookmeter.com",
		"bookmeter.com",
		"www.bookmeter.com"))

	c.SetRequestTimeout(time.Duration(35) * time.Second)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// c.OnScraped(func(r *colly.Response) {
	// 	fmt.Println("Finished", r.Request.URL)
	// })

	c.OnHTML("ul.book-list__group", func(e *colly.HTMLElement) {

		dom := e.DOM
		title := dom.Find("div.detail__title")
		fmt.Println(title.Text())
	})
	c.Visit(target)
}
