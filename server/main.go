package main

import (
	"flag"
	"fmt"

	"github.com/yunkaiyueming/netburn/g"
	_ "github.com/yunkaiyueming/netburn/server/cron"
	"github.com/yunkaiyueming/netburn/server/service"
)

func main() {
	flagParse()
	//json file parse
	//cron start
	service.StartService()
}

//flag parse
func flagParse() {
	t := flag.String("t", "", "input what you want to talk")
	flag.Parse()
	if *t != "" {
		fmt.Println("welcome talk :", *t)
	}

	args := flag.Args()
	if len(args) > 0 {
		if args[0] == "v" || args[0] == "version" {
			fmt.Println("the server version:", g.SERVER_VERSION)
		}
	}
}
