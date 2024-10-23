package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// https://mholt.github.io/json-to-go/

type Producter []struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      struct {
		Rate  float64 `json:"rate"`
		Count int     `json:"count"`
	} `json:"rating"`
}

func main() {
	url := "https://fakestoreapi.com/products"

	res, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(body))
	data_obj := Producter{}
	jsonErr := json.Unmarshal(body, &data_obj)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	for key, value := range data_obj {
		fmt.Println(key)
		fmt.Println(value.Title)
	}
}
