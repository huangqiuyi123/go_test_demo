package util

import (
	"math/rand"
	"time"
)

func RandomString(count, length int) []string {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	//const length = 10

	// 用于生成随机数的随机数生成器
	// 使用当前时间戳作为种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成指定数量的随机字符串
	result := make([]string, count)
	for i := 0; i < count; i++ {
		// 生成随机字符串
		bytes := make([]byte, length)
		for j := 0; j < length; j++ {
			bytes[j] = alphabet[r.Intn(len(alphabet))]
		}

		// 保存随机字符串到结果列表中
		result[i] = string(bytes)
	}

	return result
}
