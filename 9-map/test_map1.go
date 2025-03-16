package main

import "fmt"

func main() {
	//第一种声明方式:声明myMap1是一种map类型
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1是一个空map")
	}
	//在使用map前，需要先用make给map分配空间
	myMap1 = make(map[string]string, 10)
	myMap1["name"] = "小王子"
	myMap1["addr"] = "沙河"
	fmt.Println(myMap1)

	//第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "小王子"
	myMap2[2] = "沙河"
	fmt.Println(myMap2)

	//第三种声明方式  在初始化的时候，已经开辟了空间
	myMap3 := map[string]string{
		"name": "小王子",
		"addr": "沙河",
	}
	fmt.Println(myMap3)
}
