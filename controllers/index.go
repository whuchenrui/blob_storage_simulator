package main

import (
	"net/http"
	"github.com/whuchenrui/blob_storage/models"
)


func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/css/", fs)
	http.Handle("/js/", fs)

	http.HandleFunc("/index", models.IndexHandler)
	http.HandleFunc("/update/tc", models.TcUpdateHandler)
	http.ListenAndServe(":8080", nil)
}
