package main

import (
	"fmt"
	"io/ioutil"
)

//使用ioutil包方法进行文件拷贝
func main() {
	//声明两个已存在的文件
	file1Path := "C:\\data\\temp\\path\\abc.txt"
	file2Path := "C:\\data\\temp\\path\\ddd.txt"

	//读取文件1的内容
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("文件读取错误:%v\n", err)
		return
	}
	//将读取的内容写入文件2
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("文件写入错误:%v\n", err)
	}

}
