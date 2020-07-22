package main

import "fmt"

//只写
func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}

//只读
func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("只读管道读出的数据:", v)
	}
	var a struct{}
	exitChan <- a
}

func main() {
	var ch chan int
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	//遍历管道
	//遍历后, 仅有唯一类型数据,且不愿作处理, 即可省略不写.实际写法: for _ = range exitChan{...}
	for range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("程序结束...")
}
