package main

import (
	"fmt"
)

//select 关键字 类似switch. 是针对管道channel的
//select 可用于解决: 从管道取数据 阻塞的问题
func main() {
	//放入 10个 int数据
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//放入 5个 string数据
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//遍历管道时,必须先关闭管道,但项目中无法确定何时关闭管道
	//可以使用select
	//即使用select替代for循环
	//从输出结果可知 select循环是无序的, 并发的 取数据
	for {
		select {
		case i := <-intChan:
			fmt.Printf("从intChan管道取出的数据: %v\n", i)
			//time.Sleep(time.Second)
		case s := <-strChan:
			fmt.Printf("从strChan管道取出的数据: %v\n", s)
			//time.Sleep(time.Second)
		default:
			fmt.Println("未取到任何数据...")
			//time.Sleep(time.Second)
			return
		}
	}
}
