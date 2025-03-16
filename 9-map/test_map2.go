package main

import "fmt"

func printMap(cityMap map[string]string) {

	//引用传递 指针
	for key, value := range cityMap {
		fmt.Println("key=", key, "value=", value)
	}

}

func main() {
	cityMap := make(map[string]string)
	cityMap["北京"] = "北京"
	cityMap["上海"] = "上海"
	cityMap["深圳"] = "深圳"

	printMap(cityMap)

	delete(cityMap, "北京")
	cityMap["上海"] = "xxx"
	fmt.Println("------------------------")
	printMap(cityMap)
}
