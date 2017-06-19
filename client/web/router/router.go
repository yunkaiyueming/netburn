package router

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/yunkaiyueming/netburn/client/models/cron"
	"github.com/yunkaiyueming/netburn/client/models/users"
	"github.com/yunkaiyueming/netburn/utils"
)

func Router() {
	//lobby
	http.HandleFunc("/", HandleLobby)
	http.HandleFunc("/index", HandleLobby)
	http.HandleFunc("/lobby", HandleLobby)

	//user
	http.HandleFunc("/user/all_user", HandleAllUser)
	http.HandleFunc("/user/get_user_by_email", HandleGetUserByEmail)
	http.HandleFunc("/user/upsert_user", HandleUpsertUser)

	//cron
	http.HandleFunc("/cron/list", HandleAllCron)
}

//==============Lobby===========================
func HandleLobby(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/index" && r.URL.Path != "/lobby" && r.URL.Path != "/favicon.ico" {
		Handle404(w, r)
		return
	}

	actionItems, sortModels := utils.GetActionConfig()
	dataMap := map[string]interface{}{
		"ActionItems": actionItems,
		"SortModles":  sortModels,
	}

	t, err := template.ParseFiles("./web/views/index.html")
	//t.Funcs(utils.RegisterTplFunc())
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Execute(w, dataMap)
}

//==============User===========================
func HandleAllUser(w http.ResponseWriter, r *http.Request) {
	uItems := users.DefaultUserClient.GetAllUsers()
	ret, err := json.Marshal(uItems)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(ret)
	}
}

func HandleGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	item := users.DefaultUserClient.GetUserByEmail(email)
	ret, err := json.Marshal(item)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(ret)
	}
}

func HandleUpsertUser(w http.ResponseWriter, r *http.Request) {
	item := users.DefaultUserClient.UpsertUser()
	ret, err := json.Marshal(item)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(ret)
	}
}

//==============Cron===========================
func HandleAllCron(w http.ResponseWriter, r *http.Request) {
	cItems := cron.DefaultCronClient.GetAllCrons()
	ret, err := json.Marshal(cItems)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(ret)
	}
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	fmt.Println("==404==")
	t, _ := template.ParseFiles("./web/views/404.html")
	t.Execute(w, "")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		Handle404(w, r)
	}
}
