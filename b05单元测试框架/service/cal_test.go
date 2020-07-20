package service

import "testing"

//单元测试框架
//文件后缀必须是 "_test.go"
//函数名必须是 "Test"+ 函数名
//形参必须是 "*testing.T"

//用于测试cal.go里的函数
func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	//这里模拟出错场景, 输出错误信息	t.Fatalf(): 输出错误信息,并退出程序
	//if res != 55 {
	//	t.Fatalf("addUpper函数执行错误,期望值=%v 实际值=%v\n", 55, res)
	//}
	//日志输出
	t.Log("测试正常", res)
}

func TestGetSub(t *testing.T) {
	res := getSub(10, 3)
	t.Logf("测试结果:%v", res)
}
