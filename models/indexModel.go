package models

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
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
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out[:])
	fmt.Fprint(w, s)
}
