// 7) Json

package main

import (
	"encoding/json"
	"fmt"
)

//Book is
type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

//Author is
type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

func main() {

	author := Author{Sales: 3, Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Golang", Author: author}

	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}
