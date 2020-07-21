package main

import (
	"fmt"
	"sync"
)

//计算1-200的各个数的和，并且把各个数的和放入到map中
//问题1: 报错concurrent map writes.		全局变量未加锁,发生资源争夺. 即多个协程同时并发写map
//go build -race start.go	(-race 找出程序中发生资源竞争的数据)
//解决1: 使用互斥锁 或者	[使用管道channel存储数据]

//问题2: 协程还在执行,而主线程已执行完毕,终止了程序
//解决2: 使用管道channel
var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁 lock
	//Mutex: 互斥
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		//如果这里使用阶乘,数值过大会超过int长度,则输出 res = 0
		//res *= i
		res += i
	}
	//加锁
	lock.Lock()
	myMap[n] = res //报错concurrent map writes
	//解锁
	lock.Unlock()
}

func main() {
	for i := 0; i <= 200; i++ {
		//开启协程 执行test()函数
		go test(i)
	}
	//休眠10s	等待协程执行完毕
	//time.Sleep(time.Second * 10)

	//这里加锁是因为即便程序声明等待10s,而底层代码不会理解,仍然发生资源竞争
	lock.Lock() //加锁
	//输出myMap
	for i, v := range myMap {
		fmt.Printf("myMap[%d]=%d\n", i, v)
	}
	lock.Unlock() //解锁
}
