package main

import "fmt"

func func1() {
	fmt.Println("A")
}
func func2() {
	fmt.Println("B")
}
func func3() {
	fmt.Println("C")
}
func defer_call() {
	defer func1()
	defer func2()
	defer func3()
}
func deferFunc() int {
	fmt.Println("defer func....")
	return 0
}
func returnFunc() int {
	fmt.Println("return func....")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()

	return returnFunc() //return先执行  defer后执行

}
func main() {
	defer_call()
	returnAndDefer()
}
