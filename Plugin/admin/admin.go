package admin

import "github.com/xww2652008969/wbot/client"

type Admin struct{}

func (a Admin) PluginName() string {
	return "admin"
}

func (a Admin) PluginVersion() string {
	return "1.0.0"
}

func (a Admin) PluginAuthor() string {
	return "xww"
}

func (a Admin) GroupHandle(client client.Client, message client.Message) {
	if message.UserId != 1271701079 {
		return
	}
	api := client.Newsenapi()
	m := api.Getchatmessage()
	m.Group_id = message.GroupId
	switch message.RawMessage {
	case "插件列表":
		for _, v := range client.Pluginslist {
			if v.PluginName() != a.PluginName() {
				m.AddText(v.PluginName() + " " + v.PluginVersion() + "\n")
			}
		}
		api.SendGroupMsg()
	}
}

func (a Admin) PrivateHandle(client client.Client, message client.Message) {

}

func (a Admin) MessageSendhandle(client client.Client, message client.Message) {

}

func (a Admin) NoticeHandle(client client.Client, message client.Message) {

}

func (a Admin) Push(client *client.Client) {

}
