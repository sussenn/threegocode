package main

import (
	"bufio"
	"fmt"
	"os"
)

//bufio包方法.	文件写操作 	os.OpenFile()
func main() {
	filePath := "C:\\data\\temp\\path\\abc.txt"
	//获取文件
	//O_WRONLY: 只写模式打开文件
	//O_CREATE: 如果不存在将创建一个新文件
	//O_RDONLY: 只读模式打开文件
	//O_RDWR: 读写模式打开文件
	//O_APPEND: 写操作时,将数据附件到文件末尾
	//O_EXCL: 配合O_CREATE使用, 文件必须不存在
	//O_SYNC: 打开文件用于同步I/O
	//O_TRUNC: 如果可能, 打开时清空文件
	//perm参数: 文件权限(Linux下才有效)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	str := "hello golang\r\n"
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		//返回写入数据的长度
		write.WriteString(str)
	}
	//刷新, 将缓冲区数据放入文件
	_ = write.Flush()
}
