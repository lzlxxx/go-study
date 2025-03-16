package main

import "fmt"

func swap(p *int, q *int) {
	var temp int
	temp = *p
	*p = *q
	*q = temp

}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("a=", a, "b=", b)

	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int
	pp = &p
	fmt.Println(pp)
	fmt.Println(&p)

}
