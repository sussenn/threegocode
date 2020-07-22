package main

import (
	"fmt"
	"reflect"
)

//Elem()返回Value.	即变量指向的值
func testInt(b interface{}) {
	//获取入参的值
	val := reflect.ValueOf(b)
	//确定入参的类型,然后进行设置
	v := val.Elem()
	fmt.Println("Elem() 获取入参的指向值:", v)
	v.SetInt(110)
	fmt.Printf("反射修改后值的 val本身地址值=%v\n", val)
}

func main() {
	num := 10
	//通过反射,修改了 num的值
	testInt(&num)
	fmt.Println("使用反射函数后,值变化 num=", num)
}
