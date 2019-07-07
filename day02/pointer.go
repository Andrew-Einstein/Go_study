package main

import "fmt"

// 指针的简单应用
/*
任何类型的指针变量的零值是nil，也就是空指针
p := &x 获取变量x的内存地址
p != nil
两个指针相同，两个指针指向同一个变量，或都是空值nil
*/
func main() {
	x := 12
	p := &x
	fmt.Println(x,p)
	fmt.Println(*p)

	*p = 32
	fmt.Println(x)  //为p指向的地址对应的变量赋值

	var x1,y1 int
	fmt.Println(&x1 == &x1,&x1 == &y1, &x1 == nil)
	fmt.Println(&x1,&y1)

	pp := &x1
	fmt.Println(pp == nil)
	pp = nil
	fmt.Println(pp == nil)
}
