package models

import (
	"fmt"
	"html/template"
	"net/http"
)

type data struct {
	Title string
	Body string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index", r)
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, &data{Title: "pars", Body: "World"})
}

func TcUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("view", r)

}
