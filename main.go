package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"lianjia_spider/internal/house_list"
	"lianjia_spider/pkg/db"
	"net/http"
)

const (
	url       = "https://cd.lianjia.com/ershoufang/shuangliu/sf1/"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"
	keepAlive = true
)

type PageInfo struct {
	TotalPage   int `json:"totalPage"`
	CurrentPage int `json:"curPage"`
}

func GetPageInfo(pages string) PageInfo {
	var pageInfo PageInfo
	_ = json.Unmarshal([]byte(pages), &pageInfo)
	return pageInfo
}

func main() {

	selectorHouseInfo := "#content > div.leftContent > ul"
	selectorPageInfo := "#content > div.leftContent > div.contentBottom.clear > div.page-box.fr > div"

	c := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	c.WithTransport(&http.Transport{DisableKeepAlives: keepAlive})

	c.OnHTML(selectorHouseInfo, house_list.OnHTMLFuncGetHouseInfoList("双流"))

	//c.OnHTML(selectorHouseInfo, func(e *colly.HTMLElement) {
	//	e.ForEach("li", func(i int, element *colly.HTMLElement) {
	//
	//		if i < 1 {
	//			fmt.Println("end")
	//			return
	//		}
	//		fmt.Println(i)
	//		house := db.HouseInfo{
	//			CreatedTimeDay:   time.Now().Day(),
	//			CreatedTimeMonth: time.Now().Month().String(),
	//		}
	//
	//		house_id := element.Attr("data-lj_action_housedel_id")
	//		fmt.Println("houseID", house_id)
	//
	//		house.Title = element.ChildText("div.info.clear > div.title > a")
	//		fmt.Println("title", house.Title)
	//		infoString := element.ChildText("div.info.clear > div.address > div")
	//		fmt.Println(infoString)
	//		totalPrice := element.ChildText("div.info.clear > div.priceInfo > div.totalPrice.totalPrice2 > span")
	//		fmt.Println("totalPrice", totalPrice)
	//
	//		unitPrice := element.ChildText("div.info.clear > div.priceInfo > div.unitPrice")
	//		fmt.Println("unitPrice", unitPrice)
	//
	//	})
	//})

	c.OnHTML(selectorPageInfo, func(element *colly.HTMLElement) {
		pageInfoString := element.Attr("page-data")
		pageInfo := GetPageInfo(pageInfoString)
		fmt.Println(pageInfo)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnScraped(func(response *colly.Response) {
		db.CreatHouseInfo(house_list.HouseList)
	})

	c.Visit(url)
}
