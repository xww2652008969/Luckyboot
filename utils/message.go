package utils

import (
	"github.com/xww2652008969/wbot/client"
	"strconv"
)

func Isat(m client.Message) (flag bool, qq int64) {
	for _, v := range m.Message {
		if v.Type == "at" {
			i, _ := strconv.ParseInt(v.Data.Qq, 10, 64)
			return true, i
		}
	}
	return false, 0
}

// 获取名字 如果card为空返回nickname
func Getusername(Card, Nickname string) string {
	if Card == "" {
		return Nickname
	}
	return Card
}
