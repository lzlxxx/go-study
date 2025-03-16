package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{
		Title:  "json",
		Year:   2018,
		Price:  100,
		Actors: []string{"Jack", "Tom"},
	}
	//编码的过程，结构体-》json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}
	fmt.Printf("jsonStr=%s\n", jsonStr)

	//解码的过程，json-》结构体
	//jsonStr={"title":"json","year":2018,"price":100,"actors":["Jack","Tom"]}
	my_movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil {
		fmt.Println("json unmarshal failed, err:", err)
		return
	}
	fmt.Printf("my_movie=%v\n", my_movie)

}
