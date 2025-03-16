package main

import "fmt"

func main() {
	//声明slice是一个切片，并且初始化，默认值是1,2,3 ，长度为3
	//slice1 := []int{1, 2, 3}

	//声明slice是一个切片，但是没有初始化分配空间，默认值是nil，长度为0
	var slice1 []int
	//slice1 = make([]int, 3)//开辟内存空间，初始化长度为3，默认值是0

	//声明slice是一个切片,同时给slice分配空间，3个空间，初始化值是0
	//var slice1 []int = make([]int, 3)

	//声明slice是一个切片,同时给slice分配空间，3个空间，初始化值是0,通过:=推导出slice是切片
	//slice1 := make([]int, 3)

	fmt.Printf("len=%d,slice=%v\n", len(slice1), slice1)

	//判断一个slice是否是0
	if slice1 == nil {
		fmt.Println("slice1 is nil")
	} else {
		fmt.Println("slice1 is not nil")
	}

}
