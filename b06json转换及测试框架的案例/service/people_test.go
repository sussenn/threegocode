package service

import "testing"

func TestPeople_Store(t *testing.T) {
	people := &People{
		Name:  "西门吹雪",
		Age:   18,
		Skill: "落花吹雪",
	}
	res := people.Store()
	if !res {
		t.Fatalf("函数Store()错误,期望值:%v 实际值;%v\n", true, res)
	}
	t.Logf("测试成功...")
}

func TestPeople_ReStore(t *testing.T) {
	var people = &People{}
	res := people.ReStore()
	if !res {
		t.Fatalf("函数ReStore()错误,期望值:%v 实际值;%v\n", true, res)
	}
	if people.Name != "西门吹雪" {
		t.Fatalf("函数ReStore()错误,期望值:%v 实际值;%v\n", "西门吹雪", people.Name)
	}
	t.Logf("测试成功...")
}
