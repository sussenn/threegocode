package main

import (
	"fmt"
	"runtime"
)

//从Go1.5版本开始，默认执行下面语句以便让代码并发执行，最大效率地利用 CPU
func main() {
	//获取当前系统cpu核数	//6
	num := runtime.NumCPU()
	fmt.Println("cpu核数:", num)
	//设置最大的可同时使用的CPU核数 6
	runtime.GOMAXPROCS(num)
}
