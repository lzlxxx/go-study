package main

import "fmt"

func finbonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x: //如果c科协，则该case就会进来
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	//sub go
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	//main go
	finbonacii(c, quit)

}
