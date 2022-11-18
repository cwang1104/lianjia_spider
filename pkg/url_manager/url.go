package url_manager

import "fmt"

// https://cd.lianjia.com/ershoufang/huafu1/sf1/
var baseUrl = map[string]string{
	"华府":    "https://cd.lianjia.com/ershoufang/huafu1/",
	"锦江生态带": "https://cd.lianjia.com/ershoufang/jinjiangshengtaidai/",
	"南湖":    "https://cd.lianjia.com/ershoufang/nanhu/",
	"华阳":    "https://cd.lianjia.com/ershoufang/huayang/",
	"海洋公园":  "https://cd.lianjia.com/ershoufang/huayang/",
}

func GetListUrl(area string, page int) string {
	return fmt.Sprintf("%ssf%d/", baseUrl[area], page)
}

func GetCountUrl(area string) string {
	return fmt.Sprintf("%ssf%d/", baseUrl[area], 1)
}
