package main

import "fmt"

//限制管道的大小, 协程同步读写intChan管道并不会出现问题.
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Printf("writeData放入的数据: %v\n", i)
	}
	close(intChan) //关闭管道
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData读取的数据: %v\n", v)
	}
	//完成读取操作后,向 "退出"管道 放入标识
	exitChan <- true
	close(exitChan)
}

func main() {
	//这里限制管道的大小, 协程同步读写intChan管道并不会出现问题.
	//如果只写不读,则intChan管道会阻塞!
	intChan := make(chan int, 5)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
