package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yunkaiyueming/netburn/client/service"
	"github.com/yunkaiyueming/netburn/client/web"
	"github.com/yunkaiyueming/netburn/g"
)

func flagParse() {
	flag.Parse()
	v := flag.Arg(0)
	if v == "v" || v == "version" { //go run main.go v
		fmt.Println("the client version:" + g.CLIENT_VERSION)
		os.Exit(0)
	}
}

func main() {
	flagParse()
	service.StartRpc()
	service.StartGrpc()
	web.StartWebServer()
}
