package models

import (
	"fmt"
	"html/template"
	//"log"
	"net/http"
	//"os/exec"
	"github.com/whuchenrui/blob_storage/config"
	//"strconv"
	"strconv"
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

func GetTcCommandStr() (res string){

	return
}

func TcUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("before: ", *config.Tc)
	UpdateTcConfig(r)
	fmt.Println("after config: ", *config.Tc)
	commandStr := GetTcCommandStr()
	fmt.Println(commandStr)

	//out, err := exec.Command(commandStr).Output()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//s := string(out[:])
	fmt.Fprint(w, commandStr)
}
