package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
}

//方法的排序默认是按照方法名的排序(ASCII码)
//方法0 	返回2数和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//方法2	接收值,给对象赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

//方法1 	显示对象的值
func (s Monster) Print() {
	fmt.Println("--start--")
	fmt.Println(s)
	fmt.Println("--end--")
}

//函数	通过反射进行操作struct入参
func TestStruct(a interface{}) {
	//获取struct入参的类型
	typ := reflect.TypeOf(a)
	//获取struct入参的值
	val := reflect.ValueOf(a)
	//通过入参值, 获取struct入参的类别
	kd := val.Kind()
	//如果不是struct,则退出	(场景应用: 设计框架时,要求绑定参数的类型)
	if kd != reflect.Struct {
		fmt.Println("not expect struct") //非期望结构
		return
	}
	//通过入参值,获取入参 结构体有多少个字段
	num := val.NumField()
	fmt.Printf("结构体有 %d 个字段\n", num)

	//遍历入参 结构体的所有字段
	for i := 0; i < num; i++ {
		//通过入参值 获取结构体字段的值
		field := val.Field(i)
		fmt.Printf("结构体字段[%d]值=%v\n", i, field)
		//通过type 获取结构体标签tag.	即json映射字段
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("结构体字段[%d]的tag=%v\n", i, tagVal)
		}
	}

	//通过入参值,获取入参 结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("结构体有%d个方法\n", numOfMethod)

	//通过入参值,调用入参 结构体的方法
	//方法的排序默认是按照函数名的排序(ASCII码)
	val.Method(1).Call(nil) //获取第二个方法,并调用

	//定义入参切片,存储方法需要传入的参数
	var params []reflect.Value
	//设置入参为 10和40
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	//调用方法GetSum(),传入参数
	//因为传入的参数是[]reflect.Value, 返回[]reflect.Value
	res := val.Method(0).Call(params)
	//从切片中取出结果
	res1 := res[0].Int()
	fmt.Printf("反射调用方法,结果:%v\n", res1)
}

func main() {
	var a = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	TestStruct(a)
}
