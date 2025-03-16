package main

import "fmt"

func fool(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

// 返回多个返回值，匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	return 888, 999
}

// 返回多个返回值，有形参名称
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	//r1 r2 属于f003的形参，初始化默认值是0     作用域是foo3整个函数体的空间
	fmt.Println("r1=", r1)
	fmt.Println("r2=", r2)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}
func main() {
	c := fool("abc", 666)
	fmt.Println("c=", c)

	ret1, ret2 := foo2("kkk", 9898)
	fmt.Println("ret1=", ret1, "ret2=", ret2)

}
