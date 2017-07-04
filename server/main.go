package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/yunkaiyueming/netburn/g"
	_ "github.com/yunkaiyueming/netburn/server/cron"
	"github.com/yunkaiyueming/netburn/server/service"
)

func main() {
	flagParse()
	go service.StartGrpc()
	//go trickGrpc()
	service.StartService()
}

//flag parse
func flagParse() {
	t := flag.String("t", "", "input what you want to talk") //go run main.go -t=hello
	flag.Parse()
	if *t != "" {
		fmt.Println("welcome talk :", *t)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) > 0 {
		if args[0] == "v" || args[0] == "version" {
			fmt.Println("the server version:", g.SERVER_VERSION) //go run main.go v
			os.Exit(0)
		}
	}
}

func trickGrpc() {
	t := time.NewTicker(30 * time.Second)
	for {
		<-t.C
		fmt.Println(service.GetGServerInfo())
	}
}
