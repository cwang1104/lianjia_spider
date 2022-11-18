package house_list

import (
	"encoding/json"
	"lianjia_spider/pkg/db"
)

const (

	//房源信息列表选择器
	selectorHouseInfo = "#content > div.leftContent > ul"

	//页码信息选择器
	selectorPageInfo = "#content > div.leftContent > div.contentBottom.clear > div.page-box.fr > div"
)

var HouseList []db.HouseInfo

type PageInfo struct {
	TotalPage   int `json:"totalPage"`
	CurrentPage int `json:"curPage"`
}

func GetPageInfo(pages string) PageInfo {
	var pageInfo PageInfo
	_ = json.Unmarshal([]byte(pages), &pageInfo)
	return pageInfo
}
