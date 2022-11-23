package house_count

import (
	"fmt"
	"github.com/gocolly/colly"
	"lianjia_spider/pkg/db"
	"net/http"
)

const (
	url       = "https://cd.lianjia.com/ershoufang/shuangliu/sf1/"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"
	keepAlive = true
)

func ALLCountRun() error {

	c := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	c.WithTransport(&http.Transport{DisableKeepAlives: keepAlive})

	var count db.CdCount

	c.OnHTML(selectorCDcount, OnHTMLFuncCHENGDUCount(&count))

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnScraped(func(response *colly.Response) {
		if count.AllCount != 0 {
			count.Create()
		}
	})

	return c.Visit(baseUrlCDAllCount)
}

func GetCDCountRun() error {
	c := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	c.WithTransport(&http.Transport{DisableKeepAlives: keepAlive})

	var count db.CdCount

	c.OnHTML(selectorCDcount, ONHTMLCDCount(&count))

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnScraped(func(response *colly.Response) {
		if count.AllCount != 0 {
			count.Create()
		}
	})

	return c.Visit(bashUrlCDCount)
}
