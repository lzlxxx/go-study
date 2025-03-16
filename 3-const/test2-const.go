package main

import "fmt"

//const 来定义枚举类型

// iota只能配合const()使用，只有在const有累加效果
const (
	// 可以在const()添加一个iota,每行的iota都会累加1，第一行的iota的默认值是0
	BEIJING  = 10 * iota //iota=0
	SHANGHAI             //iota=1
	SHENZHEN             //iota=2
)
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, k
)

func main() {
	//常量 只读属性
	const length int = 10
	const (
		x = 11
		y = 22
	)
	fmt.Println("length=", length)
}
