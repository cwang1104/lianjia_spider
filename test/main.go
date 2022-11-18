package main

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	unChineseAnd = "[\\u4e00-\\u9fa5]|[(]|[)]"
	unChinese    = "[^\\u4e00-\\u9fa5]"
)

func main() {
	text := "3室2厅 | 80.7平米 | 东北 | 简装 | 中楼层(共20层) | 2017年建 | 塔楼"

	s := strings.ReplaceAll(text, " ", "")
	fmt.Println(s)

	sls := strings.Split(s, "|")
	fmt.Println(sls)

	//area := strings.TrimSuffix(sls[1], "平米")
	chiReg := regexp.MustCompile("[\u4e00-\u9fa5]")
	area := chiReg.ReplaceAllString(sls[1], "")
	fmt.Println("面积", area)
	kuohaoChi := regexp.MustCompile("[\u4e00-\u9fa5]|[(]|[)]")
	floorCount := kuohaoChi.ReplaceAllString(sls[4], "")

	fmt.Println("总高", floorCount)

}
