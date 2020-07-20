package service

//待测试的函数.	(在cal_test.go进行测试)
func addUpper(n int) int {
	res := 0
	for i := 0; i <= n-1; i++ {
		res += 1
	}
	return res
}

func getSub(n1, n2 int) int {
	return n1 - n2
}
