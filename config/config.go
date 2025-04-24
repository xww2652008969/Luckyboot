package config

import (
	"encoding/json"
	"errors"
	"github.com/xww2652008969/wbot/client"
	"github.com/xww2652008969/wbot/client/utils"
)

var GroupHandle = make([]client.Event, 0)
var PrivateHandle = make([]client.Event, 0)
var NoticeHandle = make([]client.Event, 0)
var MessageSenthandle = make([]client.Event, 0)
var Push = make([]client.Push, 0)

type config struct {
	Wsurl      string `json:"wsurl"`
	Wspost     string `json:"wspost"`
	Wsheader   string `json:"wsheader"`
	Clienthttp string `json:"clienthttp"`
}

func Read() (config, error) {
	var c config
	data := utils.Readfile("config.json")
	if len(data) > 0 {
		json.Unmarshal(data, &c)
		return c, nil
	}
	var d config
	da, _ := json.MarshalIndent(d, "", "  ")
	utils.Writefile("config.json", da)
	return c, errors.New("第一次启动需要修改配置文件")
}
