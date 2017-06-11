package web

import (
	// "html/template"
	"log"
	"net/http"

	"github.com/yunkaiyueming/netburn/client/web/router"
)

func init() {
	StartWebServer()
}

func StartWebServer() {
	router.Router()
	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
