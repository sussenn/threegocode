package main

import (
	"encoding/json"
	"fmt"
)

//interface{} 空接口即Java object
type Monster struct {
	Name     string `json:"user_name"` //将Name字段映射成 user_name输出
	Age      int    `json:"user_age"`
	Birthday string
	Sal      float64
	Skill    string
}

//对象序列化函数
func testStruct() {
	monster := Monster{
		Name:     "叶孤城",
		Age:      18,
		Birthday: "2020-07-20",
		Sal:      8000.0,
		Skill:    "万剑归宗",
	}
	//序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误:", err)
	}
	fmt.Println("Monster对象序列化:", string(data))
}

//map序列化函数
func testMap() {
	var mapp map[string]interface{}
	mapp = make(map[string]interface{})

	mapp["name"] = "西门吹雪"
	mapp["age"] = 28
	mapp["address"] = "襄阳"

	//序列化
	data, err := json.Marshal(mapp)
	if err != nil {
		fmt.Println("序列化错误:", err)
	}
	fmt.Println("Map序列化:", string(data))
}

//切片序列化
func testSlice() {
	var slice []map[string]interface{}
	var map1 map[string]interface{}
	var map2 map[string]interface{}

	map1 = make(map[string]interface{})
	map1["name"] = "独孤求败"
	map1["age"] = 38
	map1["address"] = "光明顶"
	slice = append(slice, map1)

	map2 = make(map[string]interface{})
	map2["name"] = "东方不败"
	map2["age"] = 48
	map2["address"] = [2]string{"灵鹫宫", "东瀛"}
	slice = append(slice, map2)

	//序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化错误:", err)
	}
	fmt.Println("切片序列化:", string(data))
}

func main() {
	testMap()
	testSlice()
	testStruct()
}
