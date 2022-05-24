package main

import "fmt"

func MapCURD() {
	m := make(map[string]string, 10)
	m["apple"] = "red"     // 添加
	m["apple"] = "green"   // 修改
	delete(m, "apple")     // 删除
	v, exist := m["apple"] // 查询
	if exist {
		fmt.Printf("apple-%s\n", v)
	}
}


