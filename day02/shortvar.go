package main

import "fmt"

//简短声明变量
/*
简短变量必须在函数内部声明
x := 10
*/

//简短变量与声明的变量的区别
//var声明变量可以不初始化
func main() {
	x := 10
	y := 0.24
	fmt.Printf("x = %d,y = %f",x,y)
	a,b,c := 1,2,"abc"
	fmt.Println( a,b,c)
}
