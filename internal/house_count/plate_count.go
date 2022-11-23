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

var PlateUrls = map[string][]string{}

func GetPlateCountRun(area, plate, url string) error {
	c := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	c.WithTransport(&http.Transport{DisableKeepAlives: keepAlive})

	var dbCount db.Count
	c.OnHTML(PlateCountSelector, OnHtmlPlateCount(&dbCount, area, plate))

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

func OnHtmlPlateCount(count *db.Count, area, plate string) colly.HTMLCallback {
	return func(e *colly.HTMLElement) {

		count.Area = area
		count.Plate = plate
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
