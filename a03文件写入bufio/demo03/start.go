package main

import (
	"bufio"
	"fmt"
	"os"
)

//bufio包方法.	写入原文件,追加写入
func main() {
	filePath := "C:\\data\\temp\\path\\abc.txt"
	//获取文件
	//O_WRONLY: 只写模式打开文件
	//O_APPEND: 写操作时,将数据附件到文件末尾
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	str := "追加内容...\r\n"
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		//返回写入数据的长度
		write.WriteString(str)
	}
	//刷新, 将缓冲区数据放入文件
	_ = write.Flush()
}
