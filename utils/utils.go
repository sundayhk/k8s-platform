package utils

import "encoding/json"

func StructToMap(s interface{}) map[string]string {
	j, _ := json.Marshal(s)      // 编码为json 字节切片
	m := make(map[string]string) // 用于存储后续键值对
	json.Unmarshal(j, &m)        // json转map
	return m
}
