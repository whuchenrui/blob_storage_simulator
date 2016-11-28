package main

import (
	"net/http"
	"github.com/whuchenrui/blob_storage/models"
)


func main() {
	http.HandleFunc("/index", models.IndexHandler)
	http.HandleFunc("/view", models.ViewHandler)
	http.ListenAndServe(":8080", nil)
}
