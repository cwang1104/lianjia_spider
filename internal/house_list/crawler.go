package house_list

import (
	"fmt"
	"github.com/gocolly/colly"
	"lianjia_spider/pkg/db"
	"lianjia_spider/pkg/util"
	"strconv"
	"strings"
	"time"
)

func OnHTMLFuncGetHouseInfoList(area string) colly.HTMLCallback {
	return func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, element *colly.HTMLElement) {
			if i < 1 {
				fmt.Println("end")
				return
			}
			fmt.Println(i)
			house := db.HouseInfo{
				CreatedTimeDay:   time.Now().Day(),
				CreatedTimeMonth: time.Now().Month().String(),
				Area:             area,
			}

			house.Plate = element.ChildText("div.info.clear > div.flood > div > a:nth-child(3)")
			house.HouseId = element.Attr("data-lj_action_housedel_id")
			house.Title = element.ChildText("div.info.clear > div.title > a")
			infoString := element.ChildText("div.info.clear > div.address > div")

			//text := "3室2厅 | 80.7平米 | 东北 | 简装 | 中楼层(共20层) | 2017年建 | 塔楼"
			info := strings.Split(util.DelSpace(infoString), "|")
			house.HouseType = info[0]

			acreage, err := strconv.ParseFloat(util.DelChinese(info[1]), 64)
			if err == nil {
				house.Acreage = acreage
			}

			house.Face = info[2]
			house.Floor = info[4]
			totalHeightStr := util.DelChineseAndBrackets(info[4])
			totalHeight, err := strconv.Atoi(totalHeightStr)
			if err == nil {
				house.TotalHeight = totalHeight
			}

			if strings.HasSuffix(info[5], "建") {
				house.HouseBirth = info[5]
			}

			totalPrice := element.ChildText("div.info.clear > div.priceInfo > div.totalPrice.totalPrice2 > span")
			tpr, err := strconv.ParseFloat(totalPrice, 64)
			if err == nil {
				house.TotalPrice = int64(tpr * 10000)
			}
			unitPrice := element.ChildText("div.info.clear > div.priceInfo > div.unitPrice")
			priceStr := util.GetNumberStr(unitPrice)
			price, err := strconv.Atoi(priceStr)
			if err == nil {
				house.UnitPrice = int64(price)
			}

			publish := element.ChildText("div.info.clear > div.followInfo")
			s := util.DelSpace(publish)
			sl := strings.Split(s, "/")
			if len(sl) > 1 {
				house.PublishTime = sl[1]
			}

			HouseList = append(HouseList, house)

			fmt.Printf("%+v\n", house)

		})
	}
}

func OnHTMLFuncGetPageCount() colly.HTMLCallback {
	return func(e *colly.HTMLElement) {
		pageInfoString := e.Attr("page-data")
		pageInfo := GetPageInfo(pageInfoString)
		fmt.Println(pageInfo)
	}
}
