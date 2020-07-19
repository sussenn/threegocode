package main

import (
	"fmt"
	"os"
)

//传入文件路径,即可判断文件是否存在的函数
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		//文件存在
		return true, nil
	}
	if os.IsNotExist(err) {
		//文件不存在
		return false, nil
	}
	//如果err类型是其他类型,则返回false及此err
	return false, err
}

func main() {
	file1Path := "C:\\data\\temp\\path\\aaa.txt"
	b, err := PathExists(file1Path)
	if !b {
		if err != nil {
			fmt.Println("文件不存在,发生未知错误", err)
		}
		fmt.Println("文件不存在")
	}
	fmt.Println("文件存在")
}
