package main

import (
	"fmt"
	"reflect"
)

//对普通类型的 反射
func reflectTest01(b interface{}) {
	//通过反射获取的传入的变量的type/kind/值
	//获取传入参数的类型
	rTyp := reflect.TypeOf(b)
	fmt.Println("入参类型:", rTyp)

	//获取入参的值
	rVal := reflect.ValueOf(b)
	n := 2 + rVal.Int()
	fmt.Println("n=", n)
	fmt.Printf("入参的值:%v 类型:%T\n", rVal, rVal)

	//将获取的入参 转为 object空接口类型
	iV := rVal.Interface()

	//将 object空接口类型 断言转为 需要的类型
	num := iV.(int) //转为int类型
	fmt.Println("断言强转后,iV为可用作运算的int类型. num = ", num)
}

//对结构体struct 的反射
func reflectTest02(b interface{}) {
	//通过反射获取的传入的变量的type/kind/值
	//获取传入参数的类型
	rTyp := reflect.TypeOf(b)
	fmt.Println("入参类型:", rTyp)
	//获取入参的值
	rVal := reflect.ValueOf(b)
	//将获取的入参 转为 object空接口类型
	iV := rVal.Interface()
	fmt.Printf("object空接口的值:%v object空接口的类型:%T\n", iV, iV)
}
