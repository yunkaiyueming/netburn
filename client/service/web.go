package service

import (
	"html/template"
	"log"
	"net/http"
)

func init() {
	startWebServer()
}

func startWebServer() {
	http.HandleFunc("/", lobby)
	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func lobby(w http.ResponseWriter, r *http.Request) {
	type ApiDescs struct {
		Id   int
		Url  string
		Desc string
	}

	Apis := []ApiDescs{
		{1, "www.baidu.com", "百度"},
		{2, "www.baidu.com", "google"},
		{3, "www.baidu.com", "淘宝"},
	}
	t, err := template.ParseFiles("./views/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, &Apis)
}
