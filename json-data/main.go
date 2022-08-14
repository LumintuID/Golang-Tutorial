package main

import (
	"encoding/json"
	"fmt"
)


type DataJson struct{
	FullName string `json:"Name"`
	Nilai int
}


func encode()  {
	var object = []DataJson{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}

func decode()  {
	// decode json 
	var jsonString = `[
		{"Name": "john wick", "Nilai": 27},
		{"Name": "ethan hunt", "Nilai": 32}
	]`

	// ubah jadi byte
	var jsonData = []byte(jsonString)

	// inisialisasi struct 
	var dataJson []DataJson

	var err = json.Unmarshal(jsonData, &dataJson)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range dataJson {
		fmt.Println("nama", v.FullName)
		fmt.Println("nilai", v.Nilai)
	}
}

func main() {
	encode()
	decode()
}