package main

import "fmt"

func printArray1(myArray []int) {
	//引用传递 指针
	//_表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value=", value)
	}

}

func main() {
	myArray := []int{1, 2, 3, 4} //动态数组，切片slice
	fmt.Printf("myArray types=%T\n", myArray)
	printArray1(myArray)
}
