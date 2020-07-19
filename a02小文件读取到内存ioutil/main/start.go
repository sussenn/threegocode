package main

import (
	"fmt"
	"io/ioutil"
)

//将文件流读取到内存.	适合小文件
func main() {
	fileName := "C:\\data\\temp\\path\\pathText.txt"
	context, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("文件读取失败:", err)
	}
	//fmt.Printf("%v",context)	//输出字节
	fmt.Printf("%v", string(context))
	//ioutil.ReadFile函数的内部封装了open和close
}
