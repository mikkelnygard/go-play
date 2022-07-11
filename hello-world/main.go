package main

import (
	"fmt"
)

type Words struct {
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
}

func get_string() string {
	Input := [3]string{"tmp/main.go", "Hello", "World"}
	
	var ret string

	for _, arg := range Input[1:] {
		ret += arg+" "
	}

	return ret
}

func main() {
	str := get_string()
	fmt.Println(str)
}
