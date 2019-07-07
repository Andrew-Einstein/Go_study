package main

import "fmt"

// 用new函数创建变量
// new函数是Go语言的内建函数，使用时不需要引用任何包
/*
	new(T) 创建一个T类型的匿名变量指针
	x := new(T) *x == 0
*/
func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 123
	fmt.Println(*p)

	// 普通的声明变量方式必须指定变量名，用new()方式创建变量是不需要指定变量名的
	p1 := new(int)
	p2 := new(int)
	fmt.Println(p1 == p2)

	var new int = 20
	fmt.Print(new)
}
func newInt() *int {
	return new(int)
}
func newInt1() *int {
	var value int
	return &value
}
