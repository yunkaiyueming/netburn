package users

import (
	"fmt"

	. "github.com/yunkaiyueming/netburn/client/service"
	. "github.com/yunkaiyueming/netburn/g"
	"github.com/yunkaiyueming/netburn/utils"
)

type UserClient struct{}

var DefaultUserClient = new(UserClient)

func (c *UserClient) GetAllUsers() []User {
	var reply []User
	err := RpcClient.Call("UserServer.AllInfos", nil, &reply)
	utils.LogErr("models/users/GetAllUsers", err)

	for _, v := range reply {
		fmt.Println(v)
	}
	return reply
}

func (c *UserClient) GetUserByEmail(email string) User {
	var reply User
	err := RpcClient.Call("UserServer.GetInfoByEmail", email, &reply)
	utils.LogErr("models/users/GetUserByEmail", err)

	return reply
}

func (c *UserClient) UpsertUser() []User {
	var ret bool
	newUser := map[string]interface{}{
		"email": "new@qq.com",
		"name":  "new",
		"age":   12,
	}

	updateUser := map[string]interface{}{
		"email": "tt@qq.com",
		"name":  "update",
		"age":   33,
	}

	fmt.Println("======insert=========")
	err := RpcClient.Call("UserServer.UpsertUserInfo", newUser, &ret)
	utils.LogErr("users/call/UpsertUser", err)
	c.GetAllUsers()

	fmt.Println("======update=========")
	err = RpcClient.Call("UserServer.UpsertUserInfo", updateUser, &ret)
	utils.LogErr("users/call/UpsertUser", err)
	return c.GetAllUsers()
}
