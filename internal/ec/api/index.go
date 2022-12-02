package api

import (
	"github.com/gin-gonic/gin"
	"lianjia_spider/pkg/db"
	"log"
)

type Data struct {
	Title struct {
		Text string `json:"text"`
	} `json:"title"`
	Tooltip struct {
	} `json:"tooltip"`
	Legend struct {
		Data []string `json:"data"`
	} `json:"legend"`
	XAxis struct {
		Data []string `json:"data"`
	} `json:"xAxis"`
	YAxis struct {
	} `json:"yAxis"`
	Series []struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Data []int  `json:"data"`
	} `json:"series"`
}

func GetData(c *gin.Context) {
	all_count := db.CdCount{}

	all_counts, err := all_count.GetCount()
	if err != nil {
		log.Println("get failed", err)
	}

}
