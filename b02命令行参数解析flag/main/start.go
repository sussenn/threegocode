package main

import (
	"flag"
	"fmt"
)

func main() {
	//输入命令行命令: go build -o test.exe start.go 生成test.exe
	//再执行程序test.exe: test.exe -u root -pwd root -h 127.0.0.1 -port 3306
	//用于接收命令行参数值
	var user string
	var pwd string
	var host string
	var port int

	//&user 接收用户输入
	//"u", 即 -u 这个指定参数
	//"", 默认值
	//"用户名,默认为空", 说明
	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "host", "localhost", "主机名,默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号,默认3306")
	//[关键操作]
	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}
