package main

/*
四种变量的声明方式
*/
import (
	"fmt"
)

// 声明全局变量，方法一，方法二，方法三是可以的
var x = 100
var y int = 1200

func main() {
	//方法一：声明一个变量，默认值是0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("type of a=%T\n", a)
	//方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b=", b)
	fmt.Printf("type of c=%T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb=%s,type of bb=%T\n", bb, bb)
	//方法三：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c=", c)
	fmt.Printf("type of c=%T\n", c)

	var cc = "abcd"
	fmt.Printf("cc=%s,type of cc=%T\n", cc, cc)

	//方法四：(常用方法)省去var关键字，直接自动匹配  方法四不支持全局变量，只能在函数体内使用
	e := 100 //既初始化又赋值
	fmt.Println("e=", e)
	fmt.Printf("type of e=%T\n", e)

	//声明多个变量
	var xx, yy int = 11, 22
	fmt.Println(xx, yy)
	var kk, ll = 100, "adff"
	fmt.Println(kk, ll)

	//多行多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println(vv, jj)
}
