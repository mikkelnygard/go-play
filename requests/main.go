package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Status struct {
	Page       PageObject  `json:"Page"`
	Components []Component `json:"components"`
}

type PageObject struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Component struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func main() {
	url := "https://www.githubstatus.com/api/v2/summary.json"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error making api request")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var Status Status
	json.Unmarshal([]byte(body), &Status)

	fmt.Println(Status.Page.Name)

}
