package crawl

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"regexp"
)

func Links(url string, pattern string) ([]string, error) {
	var links []string
	linkMap := make(map[string]int)

	c := colly.NewCollector()
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		urlmatch, _ := regexp.Match(pattern, []byte(e.Attr("href")))
		if urlmatch {
			_ = e.Request.Visit(e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		link := r.URL.String()
		_, exists := linkMap[link]
		if exists {
			linkMap[link] = linkMap[link] + 1
		} else {
			linkMap[link] = 1
		}
	})

	err := c.Visit(url)
	if err != nil {
		return []string{}, err
	}

	for url, _ := range linkMap {
		links = append(links, url)
	}

	fmt.Printf("Crawled %d links from '%s' filtered by '%s'.\n", len(links), url, pattern)

	return links, nil
}
