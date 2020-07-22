package main

import (
	"fmt"
	"time"
)

//err := recover(); 协程的异常捕获
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world" + fmt.Sprintf("%d", i))
	}
}

func test() {
	//延时执行
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}
	}()
	//这里写一个错误
	var myMap map[int]string
	myMap[0] = "未make()初始化"
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok = ", i)
		time.Sleep(time.Second)
	}
}
