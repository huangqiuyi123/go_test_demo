package common

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"testDemo/config"
)

type Params struct {
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	Url     string      `json:"url"`
	Method  string      `json:"method"`
	ResBody interface{} `json:"resBody"`
}

var file_locker sync.Mutex //config file locker
var configData = config.GetConfig()

// const filename = conf.FilePath.JsonPath
const filename = "testData/addActivity.json"

// 拼接URL,并替换结构体中的URL
func replaceURL(environment string, params []Params) []Params {
	// 循环结构体
	for i := range params {
		switch environment {
		case "test":
			// 拼接URL
			params[i].Url = "https://" + configData.Handle.Test + configData.Environment.Test + params[i].Url
			// 将替换后的Params结构体重新放入切片中
			params[i] = params[i]
		case "pre":
			// 拼接URL
			params[i].Url = "https://" + configData.Handle.Pre + configData.Environment.Pre + params[i].Url
			// 将替换后的Params结构体重新放入切片中
			params[i] = params[i]
		case "pro":
			// 拼接URL
			params[i].Url = "https://" + configData.Handle.Pro + configData.Environment.Pro + params[i].Url
			// 将替换后的Params结构体重新放入切片中
			params[i] = params[i]
		}

	}
	return params
}

// ReadJson 读取json数据
func ReadJson(environment string) []Params {

	var params []Params

	file_locker.Lock()
	data, err := os.ReadFile(filename) //read config file
	file_locker.Unlock()
	if err != nil {
		fmt.Println("read json file error")
		return params
	}

	//反序列化成结构体格式
	err = json.Unmarshal(data, &params)
	if err != nil {
		fmt.Println("Unmarshal 解析JSON失败：", err)
		return params
	}

	params = replaceURL(environment, params)
	fmt.Println(params)

	return params

}

//func main() {
//
//	//ReadJson("pro")
//	//conf := ReadJson()
//	//// 序列化成json格式
//	//jsonData, err := json.Marshal(conf)
//	//if err != nil {
//	//	fmt.Println("序列化失败!")
//	//}
//	////fmt.Println("======================================================")
//	//
//	//fmt.Printf("序列化成json格式结果: %v", string(jsonData))
//}
