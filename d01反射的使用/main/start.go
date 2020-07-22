package main

import (
	"fmt"
	"reflect"
)

//type和 kind 的区别	kind:类别.	如 int类别/struct类别/	type:类型	如:main.Student类型

//对普通类型的 反射
func reflectTest01(b interface{}) {
	//通过反射获取的传入的变量的type/值
	//获取传入参数的类型
	rTyp := reflect.TypeOf(b)
	fmt.Println("reflectTest01() 入参类型:", rTyp)

	//获取入参的值
	rVal := reflect.ValueOf(b)
	//数据类型要求匹配, int类型使用 .Int()
	n := 2 + rVal.Int()
	fmt.Println("接收入参,使用Int()匹配类型为int. n=", n)
	//类型:reflect.Value
	fmt.Printf("入参的值:%v 类型:%T\n", rVal, rVal)

	//将获取的入参 转为 object空接口类型
	iV := rVal.Interface()

	//将 object空接口类型 断言转为 需要的类型
	num := iV.(int) //转为int类型
	num += 1
	fmt.Println("断言强转后,iV为可用作运算的int类型. num = ", num)
}

//对结构体struct 的反射
func reflectTest02(b interface{}) {
	//通过反射获取的传入的变量的type/kind/值
	//获取传入参数的类型
	rTyp := reflect.TypeOf(b)
	fmt.Println("reflectTest02() 入参类型:", rTyp)
	//获取入参的值
	rVal := reflect.ValueOf(b)
	//将获取的入参 转为 object空接口类型
	iV := rVal.Interface()
	fmt.Printf("object空接口的值:%v object空接口的类型:%T\n", iV, iV)
}

type Student struct {
	Name string
	Age  int
}

func main() {
	num0 := 10
	reflectTest01(num0)

	//struct的测试
	stu := Student{
		Name: "tom",
		Age:  18,
	}
	reflectTest02(stu)
}
