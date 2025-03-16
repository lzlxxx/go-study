package main

import "fmt"

func main() {
	var a string
	//pair<statictype:string,value:"xxx">
	a = "xxx"
	var allType interface{}
	allType = a
	str, _ := allType.(string)
	fmt.Println(str)
}
