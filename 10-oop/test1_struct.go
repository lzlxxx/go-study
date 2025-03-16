package main

import (
	"fmt"
)

// 声明一种新的数据类型 myint ，是int的别名
type myint int

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook(b Book) {
	//传递一个book的副本
	b.auth = "666"
}
func changeBook2(b *Book) {
	//传递一个book的指针
	b.auth = "777"
}

func main() {
	var a myint = 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b Book
	b.title = "Go语言"
	b.auth = "Go"
	fmt.Printf("%v\n", b)
	changeBook(b)
	fmt.Printf("%v\n", b)
	changeBook2(&b)
	fmt.Printf("%v\n", b)

}
