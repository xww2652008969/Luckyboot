package NudgeEvent

import (
	"Lucky/utils"
	"github.com/xww2652008969/wbot/client"
	"strings"
)

var nudge []string

func init() {
	nudge = []string{
		"喂(#`O′)，戳我干什么",
		"不许戳！",
		"再这样我要叫警察叔叔啦",
		"讨厌没有边界感的人类",
		"戳牛魔戳",
		"再戳我就要戳回去啦",
		"呜......戳坏了",
		"放手啦，不给戳QAQ",
		"(。´・ω・)ん?",
		"请不要戳 >_<",
		"这里是hmmt(っ●ω●)っ",
		"啾咪~",
		"userName有什么吩咐吗",
		"ん？",
		"hmmt不在",
		"厨房有煤气灶自己拧着玩",
		"操作太快了，等会再试试吧",
		"手指该充电啦~滴滴滴电量不足警告！",
		"这是hmmt的草莓奶油城堡~禁止投喂手指",
		"再戳的话...会触发害羞の隐藏剧情哦！",
		"操作太快像在打音游呢~要掉血啦！",
		"手指该做SPA啦~再动要变成胡萝卜啦！",
		"hmmt宝宝在睡觉~不许戳醒它！",
		"这是游戏存档点~乱戳会覆盖进度哦！",
	}
}

type NudgeEventPlugin struct {
}

func (n NudgeEventPlugin) PluginName() string {
	return "NudgeEvent"
}

func (n NudgeEventPlugin) PluginVersion() string {
	return "1.0.0"
}

func (n NudgeEventPlugin) PluginAuthor() string {
	return "xww"
}

func (n NudgeEventPlugin) GroupHandle(client client.Client, message client.Message) {
}

func (n NudgeEventPlugin) PrivateHandle(client client.Client, message client.Message) {
}

func (n NudgeEventPlugin) MessageSendhandle(client client.Client, message client.Message) {
}

func (n NudgeEventPlugin) NoticeHandle(client client.Client, message client.Message) {
	if message.TargetId == message.SelfId && message.SubType == "poke" {
		sendAPI := client.Newsenapi()
		info, err := sendAPI.GetGroupMemberInfo(message.GroupId, message.UserId)
		if err != nil {
			return
		}
		username := utils.Getusername(info.Data.Card, info.Data.Nickname)
		sendAPI.Getchatmessage().AddAt(message.UserId).AddText(" ").AddText(strings.ReplaceAll(nudge[utils.Randint(0, len(nudge)-1)], "userName", username)).Group_id = message.GroupId
		sendAPI.SendGroupMsg()
	}
}

func (n NudgeEventPlugin) Push(client *client.Client) {
}
