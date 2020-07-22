package main

import "fmt"

//管道的遍历,可以使用增强for, 或者无限for.	且遍历前必须先关闭管道
func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	//关闭管道
	close(intChan)
	//管道关闭后,无法写入数据
	//但可以读取数据
	n1 := <-intChan
	fmt.Println("关闭管道后,读取的数据:", n1)
	//==============================================
	intChan2 := make(chan int, 100)
	//放入100个数据到管道
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	//[遍历管道不能使用普通for循环!] 可以使用增强for, 或者无限for
	//使用遍历前,必须先关闭管道!
	//如果对取出的数据v,没有作用, "v :=" 可以省略不写
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("关闭管道后才能遍历数据!,v=", v)
	}
}
