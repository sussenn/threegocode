package main

import (
	"fmt"
	"sync"
)

//计算1-200的各个数的和，并且把各个数的和放入到map中
//报错concurrent map writes.		全局变量未加锁,发生资源争夺
var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁 lock
	//Mutex: 互斥
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
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
	//休眠10s
	//time.Sleep(time.Second * 10)

	lock.Lock() //加锁
	//输出myMap
	for i, v := range myMap {
		fmt.Printf("myMap[%d]=%d\n", i, v)
	}
	lock.Unlock() //解锁
}
