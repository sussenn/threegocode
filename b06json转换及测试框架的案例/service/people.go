package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type People struct {
	Name  string
	Age   int
	Skill string
}

func (this *People) Store() bool {
	//将对象转json
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("json转换异常:", err)
		return false
	}
	//json字符串写入文件
	filePath := "D:\\data\\temp\\test00.txt"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("文件写入异常:", err)
		return false
	}
	return true
}

//从文件读取json字符串转对象
func (this *People) ReStore() bool {
	filePath := "D:\\data\\temp\\test00.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("文件读取异常:", err)
		return false
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("json转换异常:", err)
		return false
	}
	return true
}
