package main

import (
	"flag"
	"fmt"

	"github.com/yunkaiyueming/netburn/client/service"
	"github.com/yunkaiyueming/netburn/client/web"
	"github.com/yunkaiyueming/netburn/g"
)

func flagParse() {
	flag.Parse()
	v := flag.Arg(0)
	if v == "v" || v == "version" {
		fmt.Println("the client version:" + g.CLIENT_VERSION)
	}
}

func main() {
	flagParse()
	service.StartRpc()
	web.StartWebServer()
}
