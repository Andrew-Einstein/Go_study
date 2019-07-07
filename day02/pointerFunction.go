package main

import "fmt"

// 指针在函数中的应用
// 函数返回指针
// 指针作为函数的参数
/*
Go可以返回局部变量指针，是安全的
复制指针，为原变量创建别名
 */

func main() {
	var xp = ff()
	fmt.Println(xp)
	fmt.Println(ff() == ff())
	fmt.Println(*xp)
	// 将指针类型作为函数参数类型
	value := 10
	fun1(value)
	fmt.Println(value)
}

func ff() *int {
	v := 10
	return &v
}
func fun1(value int) int {
	value = value * 2
	return value
}
func fun2(value *int) int {
	*value = *value * 2
	return *value
}