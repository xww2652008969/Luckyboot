package config

import (
	"Lucky/utils"
	"encoding/json"
	"errors"
)

type config struct {
	Wsurl      string `json:"wsurl"`
	Wspost     string `json:"wspost"`
	Wsheader   string `json:"wsheader"`
	Clienthttp string `json:"clienthttp"`
}

func Read() (config, error) {
	var c config
	data, _ := utils.Readfile("config.json")
	if len(data) > 0 {
		json.Unmarshal(data, &c)
		return c, nil
	}
	var d config
	da, _ := json.MarshalIndent(d, "", "  ")
	utils.Writefile("config.json", da)
	return c, errors.New("第一次启动需要修改配置文件")
}
