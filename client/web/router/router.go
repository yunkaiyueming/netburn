package router

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/yunkaiyueming/netburn/client/models/cache"
	"github.com/yunkaiyueming/netburn/client/models/cron"
	"github.com/yunkaiyueming/netburn/client/models/hfile"
	"github.com/yunkaiyueming/netburn/client/models/slog"
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

	//hfile
	http.HandleFunc("/hfile/all_bucket_files", HandleAllBucketsFiles)
	http.HandleFunc("/hfile/get_bucket_files", HandleGetBucketFiles)
	http.HandleFunc("/hfile/get_bucket_names", HandleGetBucketNames)
	http.HandleFunc("/hfile/get_file_msg", HandleGetFileMsg)

	//slog
	http.HandleFunc("/slog/app_list", HandleAppList)
	http.HandleFunc("/slog/app_by_name", HandleAppByName)
	http.HandleFunc("/slog/log_by_app", HandleLogByApp)

	//cache
	http.HandleFunc("/cache/hot_data", HandleGetHotData)
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

//==============Hfile===========================
func HandleAllBucketsFiles(w http.ResponseWriter, r *http.Request) {
	rets := hfile.DefaultHfileClient.AllBucketsFiles()
	ret, err := json.Marshal(rets)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(ret)
	}
}

func HandleGetBucketFiles(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bucketName := r.FormValue("bucket_name")

	rets := hfile.DefaultHfileClient.GetBucketFiles(bucketName)
	ret, err := json.Marshal(rets)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(ret)
	}
}

func HandleGetBucketNames(w http.ResponseWriter, r *http.Request) {
	rets := hfile.DefaultHfileClient.GetBucketNames()
	ret, err := json.Marshal(rets)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(ret)
	}
}

func HandleGetFileMsg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bucketName := r.FormValue("bucket_name")
	key := r.FormValue("key")

	rets := hfile.DefaultHfileClient.GetFileMsg(bucketName, key)
	ret, err := json.Marshal(rets)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(ret)
	}
}

//==============Slog===========================
func HandleAppList(w http.ResponseWriter, r *http.Request) {
	ret := slog.DefaultSlogClient.GetApps()
	json, err := json.Marshal(ret)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(json)
	}
}

func HandleAppByName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	appName := r.FormValue("app_name")
	ret := slog.DefaultSlogClient.GetAppByName(appName)
	json, err := json.Marshal(ret)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(json)
	}
}

func HandleLogByApp(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	appName := r.FormValue("app_name")
	ret := slog.DefaultSlogClient.GetAppLog(appName)
	json, err := json.Marshal(ret)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(json)
	}
}

//==============Cache===========================
func HandleGetHotData(w http.ResponseWriter, r *http.Request) {
	ret := cache.DefaultCacheClient.GetHotData()
	json, err := json.Marshal(ret)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(json)
	}
}
