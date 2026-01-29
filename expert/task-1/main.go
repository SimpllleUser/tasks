package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	BS          string `json:"bs"`
}

func main() {
	var user User
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		errors.New(err.Error())
	}
	defer res.Body.Close()
	fmt.Println("Response status:", res.Status)

	json.NewDecoder(res.Body).Decode(&user)
	// data, _ := json.MarshalIndent(user, "", "  ")
	data, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(data))
}
