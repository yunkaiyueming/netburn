package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego/config"
	. "github.com/yunkaiyueming/netburn/g"
)

func GetServerConfig() (map[string]string, error) {
	cf, err := config.NewConfig("ini", "./conf/app.conf")
	if err != nil {
		return nil, err
	}
	return cf.GetSection("server")
}

func GetWebServerConfig() (map[string]string, error) {
	cf, err := config.NewConfig("ini", "./config/app.conf")
	if err != nil {
		return nil, err
	}
	return cf.GetSection("client")
}

func GetActionConfig() (map[string][]ActionItems, []string) {
	cf, err := config.NewConfig("ini", "./config/action.conf")
	if err != nil {
		return nil, nil
	}

	actionItems := make(map[string][]ActionItems)
	sortModels := make([]string, 0)

	modelsOn := cf.DefaultString("ModelsOn", "")
	models := strings.Split(modelsOn, "#")
	for _, modelId := range models {
		modelSection, err := cf.GetSection(modelId)
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}

		items := make([]ActionItems, 0)
		modelName := modelSection["name"]
		delete(modelSection, "name")
		for i := 1; i <= len(modelSection); i++ {
			val := modelSection[strconv.Itoa(i)]
			varArr := strings.Split(val, "#")
			item := ActionItems{strconv.Itoa(i), varArr[0], varArr[1]}
			items = append(items, item)
		}

		actionItems[modelName] = items
		sortModels = append(sortModels, modelName)
	}

	return actionItems, sortModels
}
