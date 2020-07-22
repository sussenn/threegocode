package main

import "fmt"

//管道默认是双向的
func main() {
	//双向读写管道
	//var chan0 chan int
	//只写管道
	var chan1 chan<- int
	//只读管道
	var chan2 <-chan int

	//只写管道的初始化
	chan1 = make(chan int, 3)
	chan1 <- 100
	fmt.Println("只写管道")

	//只读管道
	n1 := <-chan2
	fmt.Println("只读管道,n1 =", n1)

}
