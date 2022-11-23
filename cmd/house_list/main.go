package main

import (
	"fmt"
	"lianjia_spider/internal/house_count"
	"lianjia_spider/pkg/util"
	"strings"
	"time"
)

func main() {
	err := house_count.ALLCountRun()
	if err != nil {
		fmt.Println("err1", err)
	}
	a := util.RandomInt(5)
	time.Sleep(time.Second * time.Duration(a))

	err = house_count.GetCDCountRun()
	if err != nil {
		fmt.Println("err2", err)
	}

	a = util.RandomInt(5)
	time.Sleep(time.Second * time.Duration(a))

	err = house_count.GetAreaUrlRun()
	if err != nil {
		fmt.Println("count.run", err)
	}

	for k, v := range house_count.AreaUrls {
		if util.CheckArea(k) {
			_ = house_count.GetAreaCountRun(k, v)
			a := util.RandomInt(5)
			time.Sleep(time.Second * time.Duration(a))
			fmt.Println(k, "ok")
		}
	}

	for k, value := range house_count.PlateUrls {
		if util.CheckArea(k) {
			for _, v := range value {
				plateAndUrl := strings.Split(v, "*")
				if len(plateAndUrl) > 1 {
					plate := plateAndUrl[0]
					url := plateAndUrl[1]
					_ = house_count.GetPlateCountRun(k, plate, url)
					a := util.RandomInt(5)
					time.Sleep(time.Second * time.Duration(a))
					fmt.Println(plate, "ok")
				}
			}
		}
	}
}
