package main

import (
	"fmt"
	"os"
)

//变量
/*
var 变量名 数据类型 = 表达式
数据类型和表达式可以省略一个
var 变量名 = 表达式
var 变量名 数据类型
*/
func main() {
	var s string
	fmt.Println(s)  //默认为空字符串
	var n int
	fmt.Println(n)  //默认为0
	var i,j,k int   //可以同时给多个变量做初始化
	fmt.Println(i,j,k)
	var b,f1,e = 2,3,"four"
	fmt.Println(b,f1,e)

	var f,err = os.Open("D:/go/Go_study/day02/declaration.go")
	if err == nil {
		fmt.Println(f.Name())
	}
}
