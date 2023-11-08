package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	response, err := http.Get(url)
	checkError(err)

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	checkError(err)

	content := string(bytes)
	tours := tourFromJson(content)

	for _, tour := range tours {
		fmt.Printf("Tour: %v, Price: %v \n", tour.Name, tour.Price)
		// fmt.Printf("Description: %v\n\n", tour.Description)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func tourFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name  string
	Price string
}
