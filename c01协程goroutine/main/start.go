package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i <= 10; i++ {
		fmt.Println("hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	//go 开启一个协程
	//从控制台输出可知, test()和main主线程同时执行
	go test()

	//正常单线程执行
	for i := 0; i <= 10; i++ {
		fmt.Println("hello 大道至简" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
