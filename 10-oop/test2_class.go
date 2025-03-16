package main

import "fmt"

// 如果类名首字母大写，则该类为public类，其他类可以访问
// 如果类名首字母小写，则该类为private类，其他类不可以访问
type Hero struct {
	// 成员变量首字母大写，则该成员变量为public成员变量，其他类可以访问
	Name  string
	Ad    int
	Level int
}

//func (this Hero) Show() {
//	fmt.Println("Name =", this.Name)
//	fmt.Println("Ad =", this.Ad)
//	fmt.Println("Level =", this.Level)
//}
//func (this Hero) getName() string {
//	//this是调用该方案的一个副本 拷贝
//	return this.Name
//}
//func (this Hero) setName(newName string) {
//	this.Name = newName
//}

func (this *Hero) Show() {
	fmt.Println("Name =", this.Name)
	fmt.Println("Ad =", this.Ad)
	fmt.Println("Level =", this.Level)
}
func (this *Hero) GetName() string {
	return this.Name
}
func (this *Hero) SetName(newName string) {
	this.Name = newName
}
func main() {
	hero := Hero{Name: "zhang3", Ad: 100, Level: 10}
	hero.Show()
	hero.SetName("li4")
	hero.Show()
}
