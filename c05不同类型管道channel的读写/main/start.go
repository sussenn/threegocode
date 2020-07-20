package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	//map类型
	//初始化管道 mapChan
	mapChan := make(chan map[string]string, 5)
	//初始化 map
	map1 := make(map[string]string, 3)
	map1["name1"] = "tom"
	map1["name2"] = "叶孤城"

	map2 := make(map[string]string, 3)
	map2["add1"] = "绝情谷"
	map2["add2"] = "灵鹫宫"

	mapChan <- map1
	mapChan <- map2

	map11 := <-mapChan
	map22 := <-mapChan

	fmt.Println("取出map11=", map11, "\tmap22=", map22)
	//=============================================
	//结构体类型	struct是值类型
	//初始化管道 catChan1
	catChan1 := make(chan Cat, 10)
	cat1 := Cat{Name: "tom", Age: 11}
	cat2 := Cat{Name: "mary", Age: 5}

	catChan1 <- cat1
	catChan1 <- cat2

	cat11 := <-catChan1
	cat22 := <-catChan1
	fmt.Println("取出cat11=", cat11, "\tcat22=", cat22)
	//---------------------------------------------
	//指针类型的结构体
	catChan2 := make(chan *Cat, 10)
	catA := Cat{Name: "汤姆", Age: 19}
	catB := Cat{Name: "表哥", Age: 55}

	catChan2 <- &catA
	catChan2 <- &catB

	catAA := <-catChan2
	catBB := <-catChan2
	fmt.Println("取出catAA=", *catAA, "\tcatBB=", *catBB)
	//=======================================================
	//object类型
	objChan := make(chan interface{}, 10)
	//放入map类型
	mapz := make(map[string]string, 3)
	mapz["name1"] = "叶孤城"
	mapz["name2"] = "陆小凤"
	//放入struct类型
	catz := Cat{Name: "tom", Age: 11}

	objChan <- mapz
	objChan <- catz
	objChan <- 10

	mapzz := <-objChan
	catzz := <-objChan
	i := <-objChan
	fmt.Println("取出mapzz=", mapzz, "\tcatzz=", catzz, "i=", i)

	//断言转型取出字段
	mapzzz := mapzz.(map[string]string)
	//name1 := mapzz.(map[string]string)["name1"]	//链式写法
	fmt.Println("mapzzz[name1]=", mapzzz["name1"])
	//断言转型取出字段
	fmt.Println("catzz.Name=", catzz.(Cat).Name)
}
