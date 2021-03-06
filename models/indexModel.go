package models

import (
	"fmt"
	"github.com/whuchenrui/blob_storage_simulator/config"
	"html/template"
	//"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"bytes"
)

type data struct {
	Title string
	Body string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	req := r.Form
	fmt.Println("index", req)
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, &data{Title: "pars", Body: "World"})
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	cmd := "ping -c 1 localhost"
	fmt.Println(cmd)
	cmd = ExecCmd(cmd)
	fmt.Fprint(w, cmd)
}

func CpuHandler(w http.ResponseWriter, r *http.Request) {

}

func TcHandler(w http.ResponseWriter, r *http.Request) {
	cmd := "tc qdisc show"
	fmt.Println(cmd)
	cmd = ExecCmd(cmd)
	fmt.Fprint(w, cmd)
}

func TcUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("before: ", *config.Tc)
	UpdateTcConfig(r)
	fmt.Println("after config: ", *config.Tc)

	cmd := GetTcCmdStr()
	fmt.Println(cmd)

	cmd = ExecCmd(cmd)
	config.Tc.Init = true

	fmt.Fprint(w, cmd)
}


// helper functions
func UpdateTcConfig(r *http.Request) {
	r.ParseForm()
	tcMap := r.Form
	val, err := strconv.Atoi(tcMap["value"][0])
	if err != nil {
		fmt.Println("error doing string to int")
	}
	reqType := tcMap["type"][0]

	switch reqType{
	case "1":
		config.Tc.PkgLoss = val
	case "2":
		config.Tc.Latency = val
	case "3":
		config.Tc.Reorder = val
	}
}

func GetTcCmdStr()  string{
	var cmd string
	if config.Tc.Init {
		cmd = "sudo tc qdisc change dev lo root netem"
	} else{
		cmd = "sudo tc qdisc add dev lo root netem"
	}

	if config.Tc.PkgLoss > 0 {
		cmd += " loss " + strconv.Itoa(config.Tc.PkgLoss) + "%"
	}
	if config.Tc.Latency > 0 {
		cmd += " delay " + strconv.Itoa(config.Tc.Latency) + "ms"
	}
	if config.Tc.Reorder > 0 {
		cmd += " reorder " + strconv.Itoa(config.Tc.Reorder) + "%"
	}
	return cmd
}

func ExecCmd(cmdStr string) string {
	parts := strings.Fields(cmdStr)
	head := parts[0]
	parts = parts[1:]

	cmd := exec.Command(head, parts...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return " "
	}
	fmt.Println("Result: " + out.String())
	return out.String()
}
