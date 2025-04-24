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
	}
}
func HandleNudgeEvent() client.Event {
	return func(client client.Client, message client.Message) {
		if message.TargetId != message.SelfId && (message.SubType != "poke") {
			return
		}
		sendAPI := client.Newsenapi()
		info, err := sendAPI.GetGroupMemberInfo(message.GroupId, message.UserId)
		if err != nil {
			return
		}
		username := utils.Getusername(info.Data.Card, info.Data.Nickname)
		sendAPI.Getchatmessage().AddAt(message.UserId).AddText(strings.ReplaceAll(nudge[utils.Randint(0, len(nudge)-1)], "userName", username)).Group_id = message.GroupId
		sendAPI.SendGroupMsg()
	}
}
