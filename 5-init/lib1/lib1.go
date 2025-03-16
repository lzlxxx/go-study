package lib1

import "fmt"

// 方法名首字母大写的话该方法可以被外部调用。小写的话只能被内部调用
func Lib1Test() {
	fmt.Println("Lib1Test()....")
}

func init() {
	fmt.Println("lib1的init()方法执行....")
}
