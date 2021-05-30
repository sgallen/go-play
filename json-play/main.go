package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const endpoint = "https://jsonplaceholder.typicode.com/users"

// Easily handles specifying only a subset of the available payload fields.
type User struct {
	Id   int    `json:id`
	Name string `json:name`
}

func get(s string) []byte {
	resp, err := http.Get(s)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func main() {
	body := get(endpoint)

	var users []User
	err := json.Unmarshal(body, &users)
	if err != nil {
		panic(err)
	}

	fmt.Println(users)

	json, err := json.MarshalIndent(users, "  ", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}
