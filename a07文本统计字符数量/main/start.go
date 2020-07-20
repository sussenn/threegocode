package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int //英文个数
	NumCount   int //数字个数
	SpaceCount int //空格个数
	OtherCount int //其他字符个数
}

func main() {
	fileName := "D:\\data\\temp\\test文件.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件错误:", err)
		return
	}
	defer file.Close()
	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透,执行下一个case	//都是属于英文字符
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字符个数;%v\t数字个数:%v\t空格个数:%v\t其他字符:%v\t", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
