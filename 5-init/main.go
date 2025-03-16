package main

import (
	_ "study-go/5-init/lib1" //匿名,使用别名的之后只调用init方法，无法使用包内方法

	//mylib2 "study-go/5-init/lib2"//别名  mylib2,mylib2.Lib2Test()直接调用
	. "study-go/5-init/lib2" //调用lib2的方法可以不用. 直接方法名就行   相当于直接把lib2中的方法导入到main包下
)

func main() {
	//lib1.Lib1Test()
	//mylib2.Lib2Test()
	Lib2Test()

}
