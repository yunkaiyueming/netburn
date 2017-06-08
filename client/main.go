package main

import (
	_ "github.com/yunkaiyueming/netburn/client/service"

	"github.com/yunkaiyueming/netburn/client/models/users"
)

func main() {
	//config parse
	//start service
	users.DefaultUserClient.GetAllUsers()
	users.DefaultUserClient.GetUserByEmail("tt@qq.com")
	users.DefaultUserClient.UpsertUser()
}
