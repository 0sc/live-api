package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/kn9ts/frodo"
)

func main() {
	app := frodo.New()
	app.Get("/", home)

	app.Serve()
}

func home(w http.ResponseWriter, r *frodo.Request) {
	tmpl := template.Must(template.ParseFiles("./index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("this error: ", err)
	}
}
