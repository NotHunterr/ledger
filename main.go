package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Username        string
	Password        string
	IsAuthenticated bool
}

func main() {

	// user := User{
	// 	Username:        "Hunter",
	// 	Password:        "Chernoff",
	// 	IsAuthenticated: true,
	// }

	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/submit", submitHandler)

	log.Println("App running on port :4200 :D")
	log.Fatal(http.ListenAndServe("localhost:4200", nil))

}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
	}

	inputValue := r.Form.Get("textInput")

	if inputValue == "" {
		fmt.Fprintf(w, "Input field cannot be empty!")
		return
	}

	fmt.Fprintf(w, "Input field value: %s", inputValue)
}
