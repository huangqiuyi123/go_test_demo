package main

import (
	"log"

	"testDemo/common"
	"testDemo/config"
	"testDemo/util"
)

func main() {

	conf := config.GetConfig()
	discountCodePath := conf.FilePath.DiscountCodePath
	request := common.ReadJson("test", discountCodePath)

	// 循环发起请求json文件的数据
	for i := 0; i < len(request); i++ {
		dataType := request[i].Type
		name := request[i].Name
		url := request[i].Url
		method := request[i].Method
		resBody := request[i].ResBody

		if dataType == "coupon" {
			log.Println("执行用例：" + name)
			randomCode := util.RandomString(2, 12)
			// 替换promotionCodes字段的code数据
			newBody := util.ReplaceCode("promotionCodes", randomCode, resBody)

			if method == "POST" {
				common.Post(url, newBody, nil, nil)
			}
			if method == "GET" {
				common.Get(url, nil, nil)
			}
		}

	}
}
