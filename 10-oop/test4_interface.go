package main

import "fmt"

// 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动作的种类
}

// 具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat Sleep")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog Sleep")
}
func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color=", animal.GetColor())
	fmt.Println("type=", animal.GetType())
}
func main() {
	/*
		var animal AnimalIF //接口的数据类型 父类指针
		animal = &Cat{"yellow"}
		animal.Sleep() //调用的就是Cat的Sleep方法
		animal = &Dog{"white"}
		animal.Sleep()
	*/
	cat := Cat{"red"}
	doh := Dog{"black"}
	showAnimal(&cat)
	showAnimal(&doh)
}
