package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Words struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

func get_string() string {
	Input := os.Args

	words := Words{}
	err := json.Unmarshal([]byte(Input[1]), &words)
	if err != nil {
		log.Fatalln("Could not unmarshal json")
	}

	ret := words.FirstName + " " + words.LastName

	return ret
}

func main() {
	str := get_string()
	fmt.Println(str)
}
