package main

import (
	"bufio"
	"fmt"
	"os"
)

//bufio包方法.	清空原文件内容,重新写入
func main() {
	filePath := "C:\\data\\temp\\path\\abc.txt"
	//获取文件
	//O_WRONLY: 只写模式打开文件
	//O_TRUNC: 如果可能, 打开时清空文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	str := "清空文件内容,并重新写入...\r\n"
	write := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		//返回写入数据的长度
		write.WriteString(str)
	}
	//刷新, 将缓冲区数据放入文件
	_ = write.Flush()
}
