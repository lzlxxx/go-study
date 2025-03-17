package main

import "fmt"

/*
关闭channel：
channel不像文件一样需要经常去关闭，只有当你确定没有任何发送数据了，或者你想显式的结束range循环之类的，才回去关闭channel
关闭channel后，无法向channel再发送数据（引发panic错误后导致接收立即返回零值）
关闭channel后，可以继续从channel中接受数据
对于nil channel，无论收发都会被阻塞
*/
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
	for {
		//ok为true表示channel没有关闭，为false表示channel已经关闭
		if data, ok := <-c; ok {
			fmt.Println("num=", data)
		} else {
			break
		}
	}
	fmt.Printf("main结束")
}
