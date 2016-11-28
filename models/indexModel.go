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
	title := r.URL.Path[len("/index"):]
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, &data{Title: title, Body: "World"})
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("view", r)
	title := r.URL.Path[len("/view"):]
	t, _ := template.ParseFiles("views/view.html")
	t.Execute(w, data{Title: title, Body: "Change the world"})
}
