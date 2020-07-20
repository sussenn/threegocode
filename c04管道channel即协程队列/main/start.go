package main

import "fmt"

//管道即协程队列,保障并发协程安全.
//管道有类型限制, 即int类型管道只能存放int类型
//管道是引用类型, 即需要使用make()初始化后才能使用
func main() {
	//创建可存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//intChan本身地址:0xc000006028	指向的地址:0xc000080080
	fmt.Printf("intChan本身地址:%p\t指向的地址:%p\n", &intChan, intChan)

	//向管道写入 10
	intChan <- 10
	//向管道写入 num
	num := 211
	intChan <- num
	//向管道写入 50
	intChan <- 50
	//管道有容量限制, intChan大小为 3
	//intChan <- 11
	//管道长度:3	容量:3
	fmt.Printf("管道长度:%v\t容量:%v\n", len(intChan), cap(intChan))
	//-----------------------------------------------------------------------
	//从管道读取数据
	var n1 int
	//队列, 先进先出
	n1 = <-intChan
	//管道取出的数据:10
	fmt.Printf("管道取出的数据:%v\n", n1)
	//取出数据后管道的长度:2	容量:3
	fmt.Printf("取出数据后管道的长度:%v\t容量:%v\n", len(intChan), cap(intChan))

	n2 := <-intChan
	n3 := <-intChan
	//只有3个数据, 超出报错deadlock
	//n4 := <-intChan
	fmt.Println("n2=", n2, "n3=", n3)
}
