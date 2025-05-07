package mieyun

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/xww2652008969/wbot/client"
	"github.com/xww2652008969/wbot/client/utils"
	"io"
	"net/url"
	"strings"
	"time"
)

const apiurl = "https://api.ff14.xin/status?data_center="

var (
	pushGroup = []int64{1064163905}
	sermap    = map[string]string{
		"猫": "猫小胖",
		"猪": "莫古力",
		"狗": "豆豆柴",
		"鸟": "陆行鸟",
	}
)

type status struct {
	DataCenter      string      `json:"data_center"`
	IsUptime        bool        `json:"is_uptime"`
	LastBonusStarts []time.Time `json:"last_bonus_starts"`
	LastBonusEnds   []time.Time `json:"last_bonus_ends"`
}

func getstatus(ser string) (status, error) {
	encodedSer := url.QueryEscape(ser)
	resp, err := utils.Httpget(apiurl+encodedSer, nil)
	if err != nil {
		return status{}, fmt.Errorf("API请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return status{}, fmt.Errorf("状态码异常: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return status{}, fmt.Errorf("读取响应失败: %w", err)
	}

	var s status
	if err := json.Unmarshal(body, &s); err != nil {
		return status{}, fmt.Errorf("JSON解析失败: %w", err)
	}

	if len(s.LastBonusStarts) == 0 {
		return status{}, fmt.Errorf("无效的响应数据")
	}
	return s, nil
}

//	func formatDuration(duration time.Duration) string {
//		hours := int(duration.Hours())
//		minutes := int(duration.Minutes()) % 60
//		return fmt.Sprintf("%d小时%d分钟", hours, minutes)
//	}
//
//	func UTCtoCST(t time.Time) time.Time {
//		return t.In(time.FixedZone("CST", 8*3600))
//	}
func getRemainingTime(t time.Time) string {
	currentTime := time.Now()
	duration := currentTime.Sub(t)
	duration = 3*time.Hour - duration%(3*time.Hour)
	return fmt.Sprintf("%d小时%d分钟", int(duration.Hours()), int(duration.Minutes())%60)
}
func getCDTime(t time.Time) string {
	t = t.Add(24 * time.Hour)
	currentTime := time.Now()
	duration := t.Sub(currentTime)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	return fmt.Sprintf("%d小时%d分钟", hours, minutes)
}
func getobligedtime(t time.Time) string {
	t = t.Add(48 * time.Hour)
	currentTime := time.Now()
	duration := t.Sub(currentTime)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	return fmt.Sprintf("%d小时%d分钟", hours, minutes)
}
func PushMie() client.Push {
	return func(client client.Client) {
		c := cron.New()
		c.AddFunc("@hourly", func() {
			pushhanlde(client)
		})
		c.Start()
		select {}
	}
}

func GroupHanlde() client.Event {
	return func(client client.Client, message client.Message) {
		if strings.Contains(message.RawMessage, "灭云") {
			comlist := strings.Split(message.RawMessage, " ")
			if len(comlist) == 2 {
				if value, exists := sermap[comlist[1]]; exists {
					s, err := getstatus(value)
					if err != nil {
						fmt.Println(err)
						return
					}
					if s.IsUptime {
						chatapi := client.Newsenapi()
						chatapi.Getchatmessage().Group_id = message.GroupId
						chatapi.Getchatmessage().Addreply(message.MessageId).AddText(s.DataCenter + "\n" + "暗黑之云诛灭战现在是奖励时间\n").AddText(fmt.Sprintf("预计剩余时间%s\n", getRemainingTime(s.LastBonusStarts[0]))).AddText("发送指令 灭云 服务器(猫,鸟,猪，狗) 即可获取对应服务器的时间")
						chatapi.SendGroupMsg()
						return
					}
					chatapi := client.Newsenapi()
					chatmessage := chatapi.Getchatmessage()
					chatmessage.Group_id = message.GroupId
					chatmessage.Addreply(message.MessageId)
					chatmessage.AddText("果咩\n现在没有开启呢\n")
					chatmessage.AddText("帮你算一下吧")
					chatmessage.AddText(fmt.Sprintf("冷却结束：%s\n强制开启:%s\n", getCDTime(s.LastBonusEnds[0]), getobligedtime(s.LastBonusEnds[0])))
					chatapi.SendGroupMsg()
				}
			}
		}
	}
}
func pushhanlde(c client.Client) {
	s, err := getstatus(sermap["猫"])
	if err != nil {
		fmt.Println(err)
		return
	}
	if !s.IsUptime {
		return
	}
	chatapi := c.Newsenapi()
	chatapi.Getchatmessage().AddText("主动推送").AddText(s.DataCenter + "\n" + "暗黑之云诛灭战现在是奖励时间\n").AddText("发送指令 灭云 服务器(猫,鸟,猪，狗) 即可获取对应服务器的")
	for _, v := range pushGroup {
		chatapi.Getchatmessage().Group_id = v
		chatapi.SendGroupMsg()
	}
}
