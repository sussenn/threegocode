package main

import (
	"fmt"
	"os"
)

func main() {
	//1. 先在cmd终端 执行命令: go build -o test.exe start.go	即运行当前代码生成test.exe程序
	//2. 执行命令: test.exe aaa 123 c:/aaa/sss.txt	即运行test.exe和其后面的参数
	//3. 因为运行的test.exe实际就是当前以下代码, 所以会抓取到cmd终端输入的内容
	fmt.Println("命令行参数个数:", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}
