package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yunkaiyueming/netburn/client/web/router"
	"github.com/yunkaiyueming/netburn/utils"
)

func StartWebServer() {
	router.Router()
	fmt.Println("web server start ===> ", getWebAddress())
	err := http.ListenAndServe(getWebAddress(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getWebAddress() string {
	serverConfig, err := utils.GetWebServerConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	return serverConfig["host"] + ":" + serverConfig["port"]
}
