package house_count

import (
	"fmt"
	"github.com/gocolly/colly"
	"lianjia_spider/pkg/db"
	"lianjia_spider/pkg/util"
	"net/http"
	"strconv"
	"time"
)

var AreaUrls = map[string]string{}

func GetAreaUrlRun() error {
	c := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	c.WithTransport(&http.Transport{DisableKeepAlives: keepAlive})

	c.OnHTML(selectorArea, GetAreaUrls())

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnScraped(func(response *colly.Response) {
		fmt.Printf("%+v\n", AreaUrls)
	})

	return c.Visit(bashUrlCDCount)
}

func GetAreaUrls() colly.HTMLCallback {
	return func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, element *colly.HTMLElement) {
			url := baseUrl + element.Attr("href")
			fmt.Println("区域" + element.Text)
			AreaUrls[element.Text] = url
		})
	}
}

func GetAreaCountRun(area, url string) error {
	c := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	c.WithTransport(&http.Transport{DisableKeepAlives: keepAlive})

	var dbCount db.Count
	c.OnHTML(AreaCountSelector, OnHtmlAreaCount(&dbCount, area))

	c.OnHTML(selectorPlate, OnHtmlGetPlateUrls(area))

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnScraped(func(response *colly.Response) {
		if dbCount.Count != 0 {
			dbCount.Create()
		}
	})

	return c.Visit(url)
}

func OnHtmlAreaCount(count *db.Count, area string) colly.HTMLCallback {
	return func(e *colly.HTMLElement) {

		count.Area = area
		a := e.Text
		fmt.Println("text2", a)
		str := util.GetNumberStr(a)
		counts, _ := strconv.Atoi(str)
		count.Count = int64(counts)

		tim := time.Now()
		count.Year = int64(tim.Year())
		count.Month = tim.Month().String()
		count.Day = int64(tim.Day())

		fmt.Printf("%+v\n", count)
	}
}

func OnHtmlGetPlateUrls(area string) colly.HTMLCallback {
	var urls []string
	return func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, element *colly.HTMLElement) {
			url := baseUrl + element.Attr("href")
			plateName := element.Text
			urlStr := plateName + "*" + url
			urls = append(urls, urlStr)
		})
		PlateUrls[area] = urls
	}
}
