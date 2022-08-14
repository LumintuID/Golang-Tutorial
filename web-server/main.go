package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "selamat datang")
}

func main()  {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		var data = map[string]string{
			"Name": "Galuh",
			"Message" : "Halo selamat datang",
		}

		var t, err = template.ParseFiles("template/index.html")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})


	http.HandleFunc("/index", index)

	fmt.Println("startting server with port 3001")
	http.ListenAndServe(":3001", nil)
}