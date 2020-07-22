package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	//遍历前必须先关闭管道, 所以放完数据即关闭管道
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		time.Sleep(time.Millisecond * 5)
		num, ok := <-intChan
		if !ok { //intChan取不到数据
			break
		}
		flag = true
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num //是素数,放入管道
		}
	}
	fmt.Println("有一个primeNum协程因为取不到数据,退出")
	//还不能关闭primeChan管道....
	//"退出"管道 写入标识
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	//开启一个协程,向intChan放入8000个数据
	go putNum(intChan)
	//开启4个协程,取数据
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//这是主线程
	go func() {
		for i := 0; i < 4; i++ {
			//将"退出"管道的标识 取出丢弃
			<-exitChan
		}
		//当从exitChan取出4个结果时,就可以放心的关闭prprimeCha
		close(primeChan)
	}()
	//遍历存放素数的管道primeChan, 取出结果
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%v\n", res)
	}
	fmt.Println("main 线程结束...")
}
