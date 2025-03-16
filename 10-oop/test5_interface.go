package main

import "fmt"

// interface是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println("arg = ", arg)

	//interface{}该如何区分 此时引用的底层数据类型到底是什么？
	//给interface{}提供“类型断言”的机制
	value, ok := arg.(string)
	if ok == true {
		fmt.Println("arg is string type, value = ", value)
	} else {
		fmt.Println("arg is not string type")
		fmt.Printf("value type = %T\n", value)
	}
}

type Book1 struct {
	auth string
}

func main() {
	book := Book1{auth: "Go"}
	myFunc(book)
	myFunc(100)
}
