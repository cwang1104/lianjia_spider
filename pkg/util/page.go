package util

import "encoding/json"

type PageInfo struct {
	TotalPage   int `json:"totalPage"`
	CurrentPage int `json:"curPage"`
}

func GetPageInfo(pages string) PageInfo {
	var pageInfo PageInfo
	_ = json.Unmarshal([]byte(pages), &pageInfo)
	return pageInfo
}
