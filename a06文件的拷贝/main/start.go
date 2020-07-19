package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//文件的拷贝
//传入目标文件和原文件
func CopyFile(dstFileNmae string, srcFileName string) (written int64, err error) {
	//获取原文件
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("打开文件失败:", err)
	}
	defer srcFile.Close()
	//获取输入流read
	reader := bufio.NewReader(srcFile)

	//获取目标文件	若文件不存在则创建
	dstFile, err := os.OpenFile(dstFileNmae, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer dstFile.Close()
	//获取输出流writer
	writer := bufio.NewWriter(dstFile)

	//复制
	return io.Copy(writer, reader)
}

func main() {
	srcFileName := "C:\\Users\\susse\\Downloads\\data\\temp\\01\\320259431.mp4"
	dstFileName := "C:\\Users\\susse\\Downloads\\data\\temp\\aaa.mp4"
	_, err := CopyFile(dstFileName, srcFileName)
	if err != nil {
		fmt.Println("文件拷贝失败:", err)
	}
	fmt.Println("文件拷贝成功")
}
