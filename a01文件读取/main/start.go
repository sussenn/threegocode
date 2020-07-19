package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//文件files是指针类型.
func main() {
	//获取文件
	file, err := os.Open("C:\\data\\temp\\path\\pathText.txt")
	if err != nil {
		fmt.Println("打开文件失败:", err)
	}
	fmt.Println("打开文件:", *file)
	//关闭文件
	//err = file.Close()
	//if err != nil {
	//	fmt.Println("关闭文件失败:", err)
	//}
	defer file.Close() //延时关闭.	即本方法执行完最后一步执行此代码

	//获取输出流 read
	//缓冲区默认大小就是4096: const defaultBufSize = 4096
	reader := bufio.NewReader(file)
	//循环读
	str := ""
	for {
		//每读到一个换行,结束当前读取. 进行下一行读取
		str, err = reader.ReadString('\n')
		if err == io.EOF { //io.EOF表示文件的末尾
			break
		}
		//注意: 如果文件内容只有一行, for 循环的break先执行,那么以下内容不输出
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}
