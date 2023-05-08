package common

import (
	"fmt"
	"log"
)

func RunPreJCase(environment, filePath string) map[string]any {

	codes := make(map[string]any)
	request := ReadJson(environment, filePath)
	// 循环发起请求json文件的数据
	for i := 0; i < len(request); i++ {

		dataType := request[i].Type
		name := request[i].Name
		url := request[i].Url
		method := request[i].Method
		resBody := request[i].ResBody
		log.Println("执行用例：" + name)

		// 判断是否是前置用例
		if dataType == "pre" {

			if method == "POST" {
				response, err := Post(url, resBody, nil, nil)
				if err != nil {
					log.Println("执行失败：", err)
				}
				codes = response
			}

			if method == "GET" {
				response, err := Get(url, nil, nil)
				if err != nil {
					log.Println("执行失败：", err)
				}
				codes = response
			}
		}
	}
	fmt.Println(codes)
	return codes
}

//func main() {
//	path := config.GetConfig().FilePath.PreCasePath
//	runPreJCase("test", path)
//}
