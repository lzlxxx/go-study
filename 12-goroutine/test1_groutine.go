package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine :i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//创建一个新的Goroutine去执行newTask函数
	go newTask()

	fmt.Println("main end")
	//i := 0
	//for {
	//	i++
	//	fmt.Printf("main goroutine:i=%d\n", i)
	//	time.Sleep(1 * time.Second)
	//}
}
