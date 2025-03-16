package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human eat....")
}
func (this *Human) Walk() {
	fmt.Println("Human walk....")
}

type SuperMan struct {
	Human //SuperMan类继承了Human类的方法
	level int
}

// 重定义Human类中的方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan eat....")
}

// 子类新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan fly....")
}
func (this *SuperMan) Print() {
	fmt.Println("name =", this.name, "sex =", this.sex, "level =", this.level)
}
func main() {
	h := Human{"zhang3", "female"}
	h.Eat()
	h.Walk()

	fmt.Println("--------------------------")
	//定义一个子类对象
	//s:= SuperMan{Human{"li4", "male"}, 100}
	var s SuperMan
	s.name = "li4"
	s.sex = "male"
	s.level = 100
	s.Walk() //父类的方法
	s.Eat()  //子类的方法
	s.Fly()  //子类新方法
	s.Print()
}
