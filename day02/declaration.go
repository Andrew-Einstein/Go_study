package main

import "fmt"

//声明
/*
Go语言中有四种声明
1、var 变量
2、const 常量
3、type 类型
4、func 函数（功能）
*/
type ABC int
func main() {
	var x = 20
	const y = 30
	fmt.Println(x)
	fmt.Println(y)
	var a ABC = 23
	fmt.Print(a)
}
