package crawler

import (
	"github.com/gocolly/colly"
	"net/http"
)

type Crawler struct {
	c         *colly.Collector
	userAgent string
}

func NewCrawler(us string) *Crawler {
	c := colly.NewCollector(
		colly.UserAgent(us),
	)
	c.WithTransport(&http.Transport{DisableKeepAlives: true})

	return &Crawler{
		c:         c,
		userAgent: us,
	}
}

func (c *Crawler) SetOnHTML(s string, f colly.HTMLCallback) {
	c.c.OnHTML(s, f)
}

func (c *Crawler) SetOnRequest(f colly.RequestCallback) {
	c.c.OnRequest(f)
}

func (c *Crawler) StartCrawler(url string) error {
	return c.c.Visit(url)
}
