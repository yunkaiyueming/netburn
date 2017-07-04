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
	fmt.Println("rpc server start ===> ", address)

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	utils.LogErr("%s,%s", err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.LogErr("%s,%s", err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Printf("%s connected %s \n", conn.RemoteAddr().String(), conn.LocalAddr().String())
		go jsonrpc.ServeConn(conn)
	}
}
