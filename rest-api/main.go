package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// STRUCT FAKE DATA USER
type student struct {
	ID string
	Name string
	Grade int
}

var data = []student{
	{"E001", "ethan", 21},
	{"E002", "wick", 22},
	{"E003", "bourne", 23},
	{"E004", "bond", 23},
}


// method
func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // SET HEADER

	if r.Method == "GET" {
		var result, err = json.Marshal(data) // ENCODE DATA TO JSON

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		id := r.FormValue("id") // GET QUERY PARAMS
		var res []byte
		var err error

		for _, val := range data {
			if val.ID == id {
				res, err = json.Marshal(val)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return 
				}

				w.Write(res)
				return 
			}
		}

		http.Error(w, "user not found", http.StatusNotFound)
	}

	http.Error(w, "", http.StatusBadRequest)
}


func main(){
	http.HandleFunc("/users", users) // ROUTE GET ALL USER
	http.HandleFunc("/user", user) // ROUTE GET USER BY ID

	fmt.Println("web server runningin port 3001")

	http.ListenAndServe(":3001", nil)
}