package service

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"

	"github.com/yunkaiyueming/netburn/utils"
)

func StartService() {
	serverConfig, err := utils.GetServerConfig()
	if err != nil {
		fmt.Sprintf("start service error:%s", err.Error())
		panic(err)
	}

	address := fmt.Sprintf("%s:%s", serverConfig["host"], serverConfig["port"])
	fmt.Println(address)

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	utils.LogErr("%s,%s", err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.LogErr("%s,%s", err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
