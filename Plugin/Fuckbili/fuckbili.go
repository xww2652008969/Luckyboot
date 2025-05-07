package Fuckbili

import (
	"Lucky/utils"
	"encoding/json"
	"errors"
	"github.com/xww2652008969/wbot/client"
	"io"
	"net/http"
	"regexp"
)

type Fuckbili struct {
}

func (f Fuckbili) PluginName() string {
	return "Fuckbili"
}

func (f Fuckbili) PluginVersion() string {
	return "1.0.0"
}

func (f Fuckbili) PluginAuthor() string {
	return "Xww"
}

func (f Fuckbili) GroupHandle(client client.Client, message client.Message) {
	if message.Message[0].Type == "json" {
		var b Qqcard
		json.Unmarshal([]byte(message.Message[0].Data.Data), &b)
		if b.Meta.Detail1.Appid == "1109937557" {
			url, err := getRealBilibiliUrl(b.Meta.Detail1.Qqdocurl)
			if err == nil {
				bili, bierr := getbilidata(url)
				if bierr != nil {
					return
				}
				sendAPI := client.Newsenapi()
				sendAPI.Getchatmessage().Addreply(message.MessageId).AddText("去你麻麻的小程序\n").AddText(bili.Data.Pages[0].Part + "\n").AddImage(bili.Data.Pages[0].FirstFrame).AddText("https://www.bilibili.com/video/" + url).Group_id = message.GroupId
				sendAPI.SendGroupMsg()
			}
		}
	}
}

func (f Fuckbili) PrivateHandle(client client.Client, message client.Message) {
}

func (f Fuckbili) MessageSendhandle(client client.Client, message client.Message) {
}

func (f Fuckbili) NoticeHandle(client client.Client, message client.Message) {
}

func (f Fuckbili) Push(client *client.Client) {
}

func getRealBilibiliUrl(shortUrl string) (string, error) {
	// 创建一个 HTTP 客户端
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 禁止自动跟随重定向
			return http.ErrUseLastResponse
		},
	}

	// 发送 GET 请求
	resp, err := client.Get(shortUrl)
	if err != nil {
		return "未知BV号", errors.New("未知BV号") // 返回未知 BV 号
	}
	defer resp.Body.Close() // 确保连接被关闭

	// 获取重定向的 URL
	realUrl := resp.Header.Get("Location")
	if realUrl == "" {
		return "未知BV号", nil // 如果没有找到 Location 返回未知 BV 号
	}

	// 使用正则表达式提取 BV 号
	bvIdRegex := regexp.MustCompile(`BV[0-9A-Za-z]+`)
	bvIdMatch := bvIdRegex.FindString(realUrl)

	if bvIdMatch == "" {
		return "未知BV号", errors.New("未知BV号") // 如果没有找到 BV 号，返回未知 BV 号
	}

	return bvIdMatch, nil // 返回找到的 BV 号
}
func getbilidata(bid string) (Bili, error) {
	var b Bili
	response, err := utils.Httpget("https://api.bilibili.com/x/web-interface/view?bvid="+bid, nil)
	if err != nil {
		return b, err
	}
	bytes, err := io.ReadAll(response.Body)
	json.Unmarshal(bytes, &b)
	if b.Code != 0 {
		return b, errors.New("未知错误")
	}
	return b, nil
}

type Qqcard struct {
	Ver    string `json:"ver"`
	Prompt string `json:"prompt"`
	Config struct {
		Type     string `json:"type"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
		Forward  int    `json:"forward"`
		AutoSize int    `json:"autoSize"`
		Ctime    int    `json:"ctime"`
		Token    string `json:"token"`
	} `json:"config"`
	NeedShareCallBack bool   `json:"needShareCallBack"`
	App               string `json:"app"`
	View              string `json:"view"`
	Meta              struct {
		Detail1 struct {
			Appid   string `json:"appid"`
			AppType int    `json:"appType"`
			Title   string `json:"title"`
			Desc    string `json:"desc"`
			Icon    string `json:"icon"`
			Preview string `json:"preview"`
			Url     string `json:"url"`
			Scene   int    `json:"scene"`
			Host    struct {
				Uin  int    `json:"uin"`
				Nick string `json:"nick"`
			} `json:"host"`
			ShareTemplateId   string `json:"shareTemplateId"`
			ShareTemplateData struct {
			} `json:"shareTemplateData"`
			Qqdocurl       string `json:"qqdocurl"`
			ShowLittleTail string `json:"showLittleTail"`
			GamePoints     string `json:"gamePoints"`
			GamePointsUrl  string `json:"gamePointsUrl"`
			ShareOrigin    int    `json:"shareOrigin"`
		} `json:"detail_1"`
	} `json:"meta"`
}
type Bili struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Bvid      string `json:"bvid"`
		Aid       int64  `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		TidV2     int    `json:"tid_v2"`
		Tname     string `json:"tname"`
		TnameV2   string `json:"tname_v2"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		DescV2    []struct {
			RawText string `json:"raw_text"`
			Type    int    `json:"type"`
			BizId   int    `json:"biz_id"`
		} `json:"desc_v2"`
		State    int `json:"state"`
		Duration int `json:"duration"`
		Rights   struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
			IsSteinGate   int `json:"is_stein_gate"`
			Is360         int `json:"is_360"`
			NoShare       int `json:"no_share"`
			ArcPay        int `json:"arc_pay"`
			FreeWatch     int `json:"free_watch"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int64  `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
			Vt         int    `json:"vt"`
		} `json:"stat"`
		ArgueInfo struct {
			ArgueMsg  string `json:"argue_msg"`
			ArgueType int    `json:"argue_type"`
			ArgueLink string `json:"argue_link"`
		} `json:"argue_info"`
		Dynamic   string `json:"dynamic"`
		Cid       int64  `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		Premiere                interface{} `json:"premiere"`
		TeenageMode             int         `json:"teenage_mode"`
		IsChargeableSeason      bool        `json:"is_chargeable_season"`
		IsStory                 bool        `json:"is_story"`
		IsUpowerExclusive       bool        `json:"is_upower_exclusive"`
		IsUpowerPlay            bool        `json:"is_upower_play"`
		IsUpowerPreview         bool        `json:"is_upower_preview"`
		EnableVt                int         `json:"enable_vt"`
		VtDisplay               string      `json:"vt_display"`
		IsUpowerExclusiveWithQa bool        `json:"is_upower_exclusive_with_qa"`
		NoCache                 bool        `json:"no_cache"`
		Pages                   []struct {
			Cid       int64  `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
			FirstFrame string `json:"first_frame"`
			Ctime      int    `json:"ctime"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool `json:"allow_submit"`
			List        []struct {
				Id          int64  `json:"id"`
				Lan         string `json:"lan"`
				LanDoc      string `json:"lan_doc"`
				IsLock      bool   `json:"is_lock"`
				SubtitleUrl string `json:"subtitle_url"`
				Type        int    `json:"type"`
				IdStr       string `json:"id_str"`
				AiType      int    `json:"ai_type"`
				AiStatus    int    `json:"ai_status"`
				Author      struct {
					Mid            int         `json:"mid"`
					Name           string      `json:"name"`
					Sex            string      `json:"sex"`
					Face           string      `json:"face"`
					Sign           string      `json:"sign"`
					Rank           int         `json:"rank"`
					Birthday       int         `json:"birthday"`
					IsFakeAccount  int         `json:"is_fake_account"`
					IsDeleted      int         `json:"is_deleted"`
					InRegAudit     int         `json:"in_reg_audit"`
					IsSeniorMember int         `json:"is_senior_member"`
					NameRender     interface{} `json:"name_render"`
				} `json:"author"`
			} `json:"list"`
		} `json:"subtitle"`
		IsSeasonDisplay bool `json:"is_season_display"`
		UserGarb        struct {
			UrlImageAniCut string `json:"url_image_ani_cut"`
		} `json:"user_garb"`
		HonorReply struct {
		} `json:"honor_reply"`
		LikeIcon          string `json:"like_icon"`
		NeedJumpBv        bool   `json:"need_jump_bv"`
		DisableShowUpInfo bool   `json:"disable_show_up_info"`
		IsStoryPlay       int    `json:"is_story_play"`
		IsViewSelf        bool   `json:"is_view_self"`
	} `json:"data"`
}
