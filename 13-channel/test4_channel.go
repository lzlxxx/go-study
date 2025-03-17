package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子go程正在运行:", i)
		}
		close(c) //close可以关闭一个channel
	}()

	//可以使用range来迭代不断操作channel
	for data := range c {
		fmt.Println("num=", data)
	}
	fmt.Printf("main结束")
}
