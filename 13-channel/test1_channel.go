package main

import "fmt"

func main() {
	//定义一个channel
	c := make(chan int)
	go func() {
		defer fmt.Println("gorun结束....")
		fmt.Println("go run 正在运行,,,,")
		c <- 666
	}()
	num := <-c
	fmt.Println("num=", num)
	fmt.Println("main函数结束,,,,")
}
