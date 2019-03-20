package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var users []User

type User struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
}

func main() {
	loadMockData()

	// Start server.
	mux := http.NewServeMux()

	mux.HandleFunc("/users", usersHandler)

	http.ListenAndServe(":8008", mux)
}

func loadMockData() {
	data, err := ioutil.ReadFile("./MOCK_DATA.json")
	if err != nil {
		fmt.Printf("failed loading json data from file: %v", err)
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Printf("failed unmarshalling users from json data: %v", err)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(404)
		return
	}

	userJson, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	_, err = w.Write(userJson)
	if err != nil {
		panic(err)
	}
}
