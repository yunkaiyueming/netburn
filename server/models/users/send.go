package users

import (
	"errors"
	"fmt"

	. "github.com/yunkaiyueming/netburn/g"
	"github.com/yunkaiyueming/netburn/utils"
)

type UserServer struct{}

var serverConfig map[string]string
var defaultUsers = []User{
	{"zhangsan@qq.com", "zhgnsna", 18, "ok"},
	{"lisi@qq.com", "lisi", 18, "ok"},
	{"tt@qq.com", "tt", 19, "tt"},
	{"nifengfeng@qq.com", "nifengfeng", 18, "ok"},
}

func (s *UserServer) AllInfos(input interface{}, users *[]User) error {
	utils.LogCall("UserServer:AllInfos", "", "ok")

	*users = defaultUsers
	fmt.Println(users)
	return nil
}

func (s *UserServer) GetInfoByEmail(email string, u *User) error {
	for _, user := range defaultUsers {
		if user.Email == email {
			*u = user
			return nil
		}
	}

	utils.LogInfo("user/send/GetInfoByEmail", "not find email:"+email)
	return errors.New("not find email:" + email)
}

func (s *UserServer) UpsertUserInfo(data map[string]interface{}, ret *bool) error {
	if data["email"] == "" {
		*ret = false
		return errors.New("pk is not exist")
	}

	tmpUser := User{
		Email: utils.Interface2String(data["email"]),
		Name:  utils.Interface2String(data["name"]),
		Age:   utils.Interface2Float(data["age"]),
		Memo:  utils.Interface2String(data["memo"]),
	}

	*ret = true
	for k, user := range defaultUsers {
		if user.Email == data["email"] {
			defaultUsers[k] = tmpUser
			return nil
		}
	}

	defaultUsers = append(defaultUsers, tmpUser)
	return nil
}

func GetDefaultUsers() []User {
	return defaultUsers
}
