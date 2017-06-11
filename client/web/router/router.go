package router

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/yunkaiyueming/netburn/client/models/users"
)

func Router() {
	http.HandleFunc("/", HandleLobby)
	http.HandleFunc("/user/all_user", HandleAllUser)
}

func HandleAllUser(w http.ResponseWriter, r *http.Request) {
	uItems := users.DefaultUserClient.GetAllUsers()
	ret, err := json.Marshal(uItems)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(ret)
	}
}

func HandleLobby(w http.ResponseWriter, r *http.Request) {
	type ApiDescs struct {
		Id   int
		Url  string
		Desc string
	}

	Apis := []ApiDescs{
		{1, "/user/all_user", "all_user"},
		{2, "www.baidu.com", "google"},
		{3, "www.baidu.com", "淘宝"},
	}
	dataMap := map[string]interface{}{
		"Apis": Apis,
		"Name": "test",
	}

	t, err := template.ParseFiles("./web/views/index.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Execute(w, dataMap)
}
