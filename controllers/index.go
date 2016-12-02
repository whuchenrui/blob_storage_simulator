package main

import (
	"net/http"
	"github.com/whuchenrui/blob_storage_simulator/models"
	"github.com/whuchenrui/blob_storage_simulator/config"
	"fmt"
)

func Init(){
	config.Tc = &config.TrafficControl{PkgLoss: 0, Latency: 0, Reorder: 0, Init: false}
	cmd := "sudo tc qdisc delete dev lo root netem"
	res := models.ExecCmd(cmd)
	fmt.Println("init TC ", res)
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
	http.HandleFunc("/status/tc", models.TcHandler)
	http.ListenAndServe(":8080", nil)
}
