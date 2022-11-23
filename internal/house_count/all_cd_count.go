package house_count

import (
	"fmt"
	"github.com/gocolly/colly"
	"lianjia_spider/pkg/db"
	"lianjia_spider/pkg/util"
	"strconv"
	"time"
)

// 成都挂牌总量
func OnHTMLFuncCHENGDUCount(count *db.CdCount) colly.HTMLCallback {
	return func(e *colly.HTMLElement) {

		count.CountType = "all"
		a := e.Text
		str := util.GetNumberStr(a)
		counts, _ := strconv.Atoi(str)
		fmt.Println("text1", a)
		count.AllCount = int64(counts)

		tim := time.Now()
		count.Year = int64(tim.Year())
		count.Month = tim.Month().String()
		count.Day = int64(tim.Day())

		fmt.Printf("%+v\n", count)
	}
}

// 获取成都住宅挂牌
func ONHTMLCDCount(count *db.CdCount) colly.HTMLCallback {
	return func(e *colly.HTMLElement) {

		count.CountType = "dwelling"
		a := e.Text
		fmt.Println("text2", a)
		str := util.GetNumberStr(a)
		counts, _ := strconv.Atoi(str)
		count.AllCount = int64(counts)

		tim := time.Now()
		count.Year = int64(tim.Year())
		count.Month = tim.Month().String()
		count.Day = int64(tim.Day())

		fmt.Printf("%+v\n", count)
	}
}
