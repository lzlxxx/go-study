package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"` //当前属性的标签
	Sex  string `info:sex`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println(taginfo)
		fmt.Println(tagdoc)
	}
}
func main() {
	var re resume
	findTag(&re)
}
