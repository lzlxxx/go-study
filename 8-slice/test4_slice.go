package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("numbers=%v,len=%d,cap=%d\n", numbers, len(numbers), cap(numbers))
	//向numbers切片追加一个元素
	numbers = append(numbers, 1)
	fmt.Printf("numbers=%v,len=%d,cap=%d\n", numbers, len(numbers), cap(numbers))
	numbers = append(numbers, 2)
	fmt.Printf("numbers=%v,len=%d,cap=%d\n", numbers, len(numbers), cap(numbers))
	//向一个容量已经满的切片追加元素会自动扩容，容量会翻倍
	numbers[0] = 1
	numbers = append(numbers, 3)
	fmt.Printf("numbers=%v,len=%d,cap=%d\n", numbers, len(numbers), cap(numbers))

	var numbers2 = make([]int, 3)
	fmt.Printf("numbers=%v,len=%d,cap=%d\n", numbers2, len(numbers2), cap(numbers2))
	numbers2 = append(numbers2, 1)
	fmt.Printf("numbers=%v,len=%d,cap=%d\n", numbers2, len(numbers2), cap(numbers2))
}
