package main

import (
	"net/http"
	"github.com/whuchenrui/blob_storage_simulator/models"
	"github.com/whuchenrui/blob_storage_simulator/config"
)

func Init(){
	config.Tc = &config.TrafficControl{PkgLoss: 0, Latency: 0, Reorder: 0, Init: false}
}

func main() {
	Init()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/css/", fs)
	http.Handle("/js/", fs)

	http.HandleFunc("/index", models.IndexHandler)
	http.HandleFunc("/update/tc", models.TcUpdateHandler)
	http.HandleFunc("/status/ping", models.PingHandler)
	http.HandleFunc("/status/cpu", models.CpuHandler)
	http.ListenAndServe(":8080", nil)
}
