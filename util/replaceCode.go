package util

import (
	"encoding/json"
	"log"
)

func ReplaceCode(field string, ReplaceValue []string, resBody interface{}) map[string]interface{} {

	data, _ := json.Marshal(resBody)

	var jsonData map[string]interface{}

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Println("Unmarshal解析失败： ", err)
	}

	jsonData[field] = ReplaceValue

	return jsonData

}
