package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//bufio包方法.	读取文件, 然后追加写入文件
func main() {
	filePath := "C:\\data\\temp\\path\\abc.txt"
	//获取文件
	//O_RDWR: 读写模式打开文件
	//O_APPEND: 写操作时,将数据附件到文件末尾
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	//先读取文件
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //如果读到文件末尾
			break
		}
		//输出读取得内容
		fmt.Print(str)
	}

	//追加写入文件
	str := "二次追加内容...\r\n"
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		//返回写入数据的长度
		write.WriteString(str)
	}
	//刷新, 将缓冲区数据放入文件
	_ = write.Flush()
}
