package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

//json转对象
func unmarshalStruct() {
	str := "{\"Name\":\"叶孤城\",\"Age\":18,\"Birthday\":\"2020-07-20\",\"Sal\":8000,\"Skill\":\"万剑归宗\"}"
	var monster Monster
	//json转struct, 数据已封装在monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("json转对象错误:", err)
	}
	fmt.Printf("json转对象:monster=%v\nmonster.Name=%v\n", monster, monster.Name)
}

//json转Map
func unmarshalMap() {
	strMap := "{\"address\":\"襄阳\",\"age\":28,\"name\":\"西门吹雪\"}"
	var mapp map[string]interface{}
	//不需要将mapp make初始化,因为Unmarshal函数底层已做处理
	err := json.Unmarshal([]byte(strMap), &mapp)
	if err != nil {
		fmt.Println("json转Map错误:", err)
	}
	fmt.Printf("json转Map:mapp=%v\n", mapp)
}

//json转切片
func unmarshalSlice() {
	strSli := "[{\"address\":\"光明顶\",\"age\":38,\"name\":\"独孤求败\"},{\"address\":[\"灵鹫宫\",\"东瀛\"],\"age\":48,\"name\":\"东方不败\"}]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(strSli), &slice)
	if err != nil {
		fmt.Println("json转Map错误:", err)
	}
	fmt.Printf("json转切片:slice=%v\n", slice)
}

func main() {
	unmarshalMap()
	unmarshalSlice()
	unmarshalStruct()
}
