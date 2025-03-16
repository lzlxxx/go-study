package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called..")
	fmt.Printf("%v\n", this)
}

func DoFiledAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType =", inputType.Name())
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue =", inputValue)
	//通过type获取里面的字段
	//1.获取interface的reflect.type,通过filed得到NumField，进行遍历
	//2.得到每个filed，数据类型
	//3.通过filed有一个interface()方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		filed := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s:%v=%v\n", filed.Name, filed.Type, value)
	}
	//通过type 获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s:%v\n", method.Name, method.Type)
	}
}
func main() {
	user := User{"xxx", 18}
	DoFiledAndMethod(user)

}
