package wxworkbot

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func strRef(s string) *string {
	return &s
}
func float64Ref(f float64) *float64 {
	return &f
}

func TestTextTemplateCard(t *testing.T) {
	jsonString := `{
    "msgtype":"template_card",
    "template_card":{
        "card_type":"text_notice",
        "source":{
            "icon_url":"https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
            "desc":"企业微信",
            "desc_color":0
        },
        "main_title":{
            "title":"欢迎使用企业微信",
            "desc":"您的好友正在邀请您加入企业微信"
        },
        "emphasis_content":{
            "title":"100",
            "desc":"数据含义"
        },
        "quote_area":{
            "type":1,
            "url":"https://work.weixin.qq.com/?from=openApi",
            "appid":"APPID",
            "pagepath":"PAGEPATH",
            "title":"引用文本标题",
            "quote_text":"Jack：企业微信真的很好用~\nBalian：超级好的一款软件！"
        },
        "sub_title_text":"下载企业微信还能抢红包！",
        "horizontal_content_list":[
            {
                "keyname":"邀请人",
                "value":"张三"
            },
            {
                "keyname":"企微官网",
                "value":"点击访问",
                "type":1,
                "url":"https://work.weixin.qq.com/?from=openApi"
            },
            {
                "keyname":"企微下载",
                "value":"企业微信.apk",
                "type":2,
                "media_id":"MEDIAID"
            }
        ],
        "jump_list":[
            {
                "type":1,
                "url":"https://work.weixin.qq.com/?from=openApi",
                "title":"企业微信官网"
            },
            {
                "type":2,
                "appid":"APPID",
                "pagepath":"PAGEPATH",
                "title":"跳转小程序"
            }
        ],
        "card_action":{
            "type":1,
            "url":"https://work.weixin.qq.com/?from=openApi",
            "appid":"APPID",
            "pagepath":"PAGEPATH"
        }
    }
}
`
	var templateCardMsg templateCardMessage
	err := json.Unmarshal([]byte(jsonString), &templateCardMsg)
	templateCard := templateCardMsg.TemplateCard
	assert.Nil(t, err)
	assert.Equal(t, TemplateCardTypeText, templateCard.CardType)

	assert.Equal(t, "https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0", *templateCard.Source.IconUrl)
	assert.Equal(t, "企业微信", *templateCard.Source.Desc)
	assert.Equal(t, TemplateCardSourceDescColorGrey, templateCard.Source.DescColor)

	assert.Equal(t, "欢迎使用企业微信", *templateCard.MainTitle.Title)
	assert.Equal(t, "您的好友正在邀请您加入企业微信", *templateCard.MainTitle.Desc)

	assert.Equal(t, "100", *templateCard.EmphasisContent.Title)
	assert.Equal(t, "数据含义", *templateCard.EmphasisContent.Desc)

	assert.Equal(t, TemplateCardQuoteAreaTypeUrl, templateCard.QuoteArea.Type)
	assert.Equal(t, "https://work.weixin.qq.com/?from=openApi", *templateCard.QuoteArea.Url)
	assert.Equal(t, "APPID", *templateCard.QuoteArea.AppID)
	assert.Equal(t, "PAGEPATH", *templateCard.QuoteArea.PagePath)
	assert.Equal(t, "引用文本标题", *templateCard.QuoteArea.Title)
	assert.Equal(t, "Jack：企业微信真的很好用~\nBalian：超级好的一款软件！", *templateCard.QuoteArea.QuoteText)

	assert.Equal(t, "下载企业微信还能抢红包！", *templateCard.SubTitleText)

	assert.Equal(t, "邀请人", (*templateCard.HorizontalContentList)[0].KeyName)
	assert.Equal(t, "张三", *(*templateCard.HorizontalContentList)[0].Value)
	assert.Equal(t, TemplateCardHorizontalContentTypeText, (*templateCard.HorizontalContentList)[0].Type)
	assert.Nil(t, (*templateCard.HorizontalContentList)[0].Url)
	assert.Nil(t, (*templateCard.HorizontalContentList)[0].MediaID)
	assert.Nil(t, (*templateCard.HorizontalContentList)[0].UserID)

	assert.Equal(t, "企微官网", (*templateCard.HorizontalContentList)[1].KeyName)
	assert.Equal(t, "点击访问", *(*templateCard.HorizontalContentList)[1].Value)
	assert.Equal(t, TemplateCardHorizontalContentTypeUrl, (*templateCard.HorizontalContentList)[1].Type)
	assert.Equal(t, "https://work.weixin.qq.com/?from=openApi", *(*templateCard.HorizontalContentList)[1].Url)
	assert.Nil(t, (*templateCard.HorizontalContentList)[1].MediaID)
	assert.Nil(t, (*templateCard.HorizontalContentList)[1].UserID)

	assert.Equal(t, "企微下载", (*templateCard.HorizontalContentList)[2].KeyName)
	assert.Equal(t, "企业微信.apk", *(*templateCard.HorizontalContentList)[2].Value)
	assert.Equal(t, TemplateCardHorizontalContentTypeAttachment, (*templateCard.HorizontalContentList)[2].Type)
	assert.Nil(t, (*templateCard.HorizontalContentList)[2].Url)
	assert.Equal(t, "MEDIAID", *(*templateCard.HorizontalContentList)[2].MediaID)
	assert.Nil(t, (*templateCard.HorizontalContentList)[2].UserID)

	assert.Equal(t, TemplateCardJumpTypeUrl, (*templateCard.JumpList)[0].Type)
	assert.Equal(t, "https://work.weixin.qq.com/?from=openApi", *(*templateCard.JumpList)[0].Url)
	assert.Equal(t, "企业微信官网", (*templateCard.JumpList)[0].Title)
	assert.Nil(t, (*templateCard.JumpList)[0].AppID)
	assert.Nil(t, (*templateCard.JumpList)[0].PagePath)

	assert.Equal(t, TemplateCardJumpTypeMiniApp, (*templateCard.JumpList)[1].Type)
	assert.Nil(t, (*templateCard.JumpList)[1].Url)
	assert.Equal(t, "跳转小程序", (*templateCard.JumpList)[1].Title)
	assert.Equal(t, "APPID", *(*templateCard.JumpList)[1].AppID)
	assert.Equal(t, "PAGEPATH", *(*templateCard.JumpList)[1].PagePath)

	assert.Equal(t, TemplateCardActionTypeUrl, templateCard.CardAction.Type)
	assert.Equal(t, "https://work.weixin.qq.com/?from=openApi", *templateCard.CardAction.Url)
	assert.Equal(t, "APPID", *templateCard.CardAction.AppID)
	assert.Equal(t, "PAGEPATH", *templateCard.CardAction.PagePath)
}

func TestSendTextTemplateCardMessage(t *testing.T) {
	botKey := GetTestBotKey()
	if len(botKey) == 0 {
		return
	}
	bot := New(botKey)
	err := bot.Send(TemplateCard{
		CardType: TemplateCardTypeText,
		Source: &TemplateCardSource{
			IconUrl:   strRef("https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0"),
			Desc:      strRef("企业微信"),
			DescColor: TemplateCardSourceDescColorGreen,
		},
		MainTitle: TemplateCardMainTitle{
			Title: strRef("欢迎使用企业微信"),
			Desc:  strRef("您的好友正在邀请您加入企业微信"),
		},
		EmphasisContent: &TemplateCardEmphasisContent{
			Title: strRef("100"),
			Desc:  strRef("数据含义"),
		},
		QuoteArea: &TemplateCardQuoteArea{
			Type: TemplateCardQuoteAreaTypeUrl,
			Url:  strRef("https://work.weixin.qq.com/?from=openApi"),
		},
		SubTitleText: strRef("下载企业微信还能抢红包！"),
		HorizontalContentList: &[]TemplateCardHorizontalContent{
			{
				KeyName: "邀请人",
				Value:   strRef("张三"),
			},
			{
				KeyName: "企微官网",
				Value:   strRef("点击访问"),
				Type:    TemplateCardHorizontalContentTypeUrl,
				Url:     strRef("https://work.weixin.qq.com/?from=openApi"),
			},
		},
		JumpList: &[]TemplateCardJump{
			{
				Type:  TemplateCardJumpTypeUrl,
				Url:   strRef("https://work.weixin.qq.com/?from=openAp"),
				Title: "企业微信官网",
			},
		},
		CardAction: TemplateCardAction{
			Type: TemplateCardActionTypeUrl,
			Url:  strRef("https://baidu.com"),
		},
	})
	assert.Nil(t, err)
}

func TestNewsTemplateCard(t *testing.T) {
	jsonString := `
{
    "msgtype":"template_card",
    "template_card":{
        "card_type":"news_notice",
        "source":{
            "icon_url":"https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
            "desc":"企业微信",
            "desc_color":0
        },
        "main_title":{
            "title":"欢迎使用企业微信",
            "desc":"您的好友正在邀请您加入企业微信"
        },
        "card_image":{
            "url":"https://wework.qpic.cn/wwpic/354393_4zpkKXd7SrGMvfg_1629280616/0",
            "aspect_ratio":2.25
        },
        "image_text_area":{
            "type":1,
            "url":"https://work.weixin.qq.com",
            "title":"欢迎使用企业微信",
            "desc":"您的好友正在邀请您加入企业微信",
            "image_url":"https://wework.qpic.cn/wwpic/354393_4zpkKXd7SrGMvfg_1629280616/0"
        },
        "quote_area":{
            "type":1,
            "url":"https://work.weixin.qq.com/?from=openApi",
            "appid":"APPID",
            "pagepath":"PAGEPATH",
            "title":"引用文本标题",
            "quote_text":"Jack：企业微信真的很好用~\nBalian：超级好的一款软件！"
        },
        "vertical_content_list":[
            {
                "title":"惊喜红包等你来拿",
                "desc":"下载企业微信还能抢红包！"
            }
        ],
        "horizontal_content_list":[
            {
                "keyname":"邀请人",
                "value":"张三"
            },
            {
                "keyname":"企微官网",
                "value":"点击访问",
                "type":1,
                "url":"https://work.weixin.qq.com/?from=openApi"
            },
            {
                "keyname":"企微下载",
                "value":"企业微信.apk",
                "type":2,
                "media_id":"MEDIAID"
            }
        ],
        "jump_list":[
            {
                "type":1,
                "url":"https://work.weixin.qq.com/?from=openApi",
                "title":"企业微信官网"
            },
            {
                "type":2,
                "appid":"APPID",
                "pagepath":"PAGEPATH",
                "title":"跳转小程序"
            }
        ],
        "card_action":{
            "type":1,
            "url":"https://work.weixin.qq.com/?from=openApi",
            "appid":"APPID",
            "pagepath":"PAGEPATH"
        }
    }
}

`
	var templateCardMsg templateCardMessage
	err := json.Unmarshal([]byte(jsonString), &templateCardMsg)
	templateCard := templateCardMsg.TemplateCard
	assert.Nil(t, err)
	assert.Equal(t, TemplateCardTypeNews, templateCard.CardType)

	assert.Equal(t, "https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0", *templateCard.Source.IconUrl)
	assert.Equal(t, "企业微信", *templateCard.Source.Desc)
	assert.Equal(t, TemplateCardSourceDescColorGrey, templateCard.Source.DescColor)

	assert.Equal(t, "欢迎使用企业微信", *templateCard.MainTitle.Title)
	assert.Equal(t, "您的好友正在邀请您加入企业微信", *templateCard.MainTitle.Desc)

	assert.Equal(t, "https://wework.qpic.cn/wwpic/354393_4zpkKXd7SrGMvfg_1629280616/0", templateCard.CardImage.Url)
	assert.Equal(t, 2.25, *templateCard.CardImage.AspectRatio)

	assert.Equal(t, TemplateCardImageTextAreaTypeUrl, templateCard.ImageTextArea.Type)
	assert.Equal(t, "https://work.weixin.qq.com", *templateCard.ImageTextArea.Url)
	assert.Equal(t, "欢迎使用企业微信", *templateCard.ImageTextArea.Title)
	assert.Equal(t, "您的好友正在邀请您加入企业微信", *templateCard.ImageTextArea.Desc)

	assert.Equal(t, TemplateCardQuoteAreaTypeUrl, templateCard.QuoteArea.Type)
	assert.Equal(t, "https://work.weixin.qq.com/?from=openApi", *templateCard.QuoteArea.Url)
	assert.Equal(t, "APPID", *templateCard.QuoteArea.AppID)
	assert.Equal(t, "PAGEPATH", *templateCard.QuoteArea.PagePath)
	assert.Equal(t, "引用文本标题", *templateCard.QuoteArea.Title)
	assert.Equal(t, "Jack：企业微信真的很好用~\nBalian：超级好的一款软件！", *templateCard.QuoteArea.QuoteText)

	assert.Equal(t, "邀请人", (*templateCard.HorizontalContentList)[0].KeyName)
	assert.Equal(t, "张三", *(*templateCard.HorizontalContentList)[0].Value)
	assert.Equal(t, TemplateCardHorizontalContentTypeText, (*templateCard.HorizontalContentList)[0].Type)
	assert.Nil(t, (*templateCard.HorizontalContentList)[0].Url)
	assert.Nil(t, (*templateCard.HorizontalContentList)[0].MediaID)
	assert.Nil(t, (*templateCard.HorizontalContentList)[0].UserID)

	assert.Equal(t, "企微官网", (*templateCard.HorizontalContentList)[1].KeyName)
	assert.Equal(t, "点击访问", *(*templateCard.HorizontalContentList)[1].Value)
	assert.Equal(t, TemplateCardHorizontalContentTypeUrl, (*templateCard.HorizontalContentList)[1].Type)
	assert.Equal(t, "https://work.weixin.qq.com/?from=openApi", *(*templateCard.HorizontalContentList)[1].Url)
	assert.Nil(t, (*templateCard.HorizontalContentList)[1].MediaID)
	assert.Nil(t, (*templateCard.HorizontalContentList)[1].UserID)

	assert.Equal(t, "企微下载", (*templateCard.HorizontalContentList)[2].KeyName)
	assert.Equal(t, "企业微信.apk", *(*templateCard.HorizontalContentList)[2].Value)
	assert.Equal(t, TemplateCardHorizontalContentTypeAttachment, (*templateCard.HorizontalContentList)[2].Type)
	assert.Nil(t, (*templateCard.HorizontalContentList)[2].Url)
	assert.Equal(t, "MEDIAID", *(*templateCard.HorizontalContentList)[2].MediaID)
	assert.Nil(t, (*templateCard.HorizontalContentList)[2].UserID)

	assert.Equal(t, TemplateCardJumpTypeUrl, (*templateCard.JumpList)[0].Type)
	assert.Equal(t, "https://work.weixin.qq.com/?from=openApi", *(*templateCard.JumpList)[0].Url)
	assert.Equal(t, "企业微信官网", (*templateCard.JumpList)[0].Title)
	assert.Nil(t, (*templateCard.JumpList)[0].AppID)
	assert.Nil(t, (*templateCard.JumpList)[0].PagePath)

	assert.Equal(t, TemplateCardJumpTypeMiniApp, (*templateCard.JumpList)[1].Type)
	assert.Nil(t, (*templateCard.JumpList)[1].Url)
	assert.Equal(t, "跳转小程序", (*templateCard.JumpList)[1].Title)
	assert.Equal(t, "APPID", *(*templateCard.JumpList)[1].AppID)
	assert.Equal(t, "PAGEPATH", *(*templateCard.JumpList)[1].PagePath)

	assert.Equal(t, TemplateCardActionTypeUrl, templateCard.CardAction.Type)
	assert.Equal(t, "https://work.weixin.qq.com/?from=openApi", *templateCard.CardAction.Url)
	assert.Equal(t, "APPID", *templateCard.CardAction.AppID)
	assert.Equal(t, "PAGEPATH", *templateCard.CardAction.PagePath)
}

func TestSendNewsTemplateCardMessage(t *testing.T) {
	botKey := GetTestBotKey()
	if len(botKey) == 0 {
		return
	}
	bot := New(botKey)
	err := bot.Send(TemplateCard{
		CardType: TemplateCardTypeNews,
		Source: &TemplateCardSource{
			IconUrl:   strRef("https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0"),
			Desc:      strRef("企业微信"),
			DescColor: TemplateCardSourceDescColorGreen,
		},
		MainTitle: TemplateCardMainTitle{
			Title: strRef("欢迎使用企业微信"),
			Desc:  strRef("您的好友正在邀请您加入企业微信"),
		},
		CardImage: &TemplateCardImage{
			Url:         "https://wework.qpic.cn/wwpic/354393_4zpkKXd7SrGMvfg_1629280616/0",
			AspectRatio: float64Ref(2.25),
		},
		ImageTextArea: &TemplateCardImageTextArea{
			Type:     TemplateCardImageTextAreaTypeUrl,
			Url:      strRef("https://work.weixin.qq.com"),
			Title:    strRef("欢迎使用企业微信"),
			Desc:     strRef("您的好友正在邀请您加入企业微信"),
			ImageUrl: "https://wework.qpic.cn/wwpic/354393_4zpkKXd7SrGMvfg_1629280616/0",
		},
		EmphasisContent: &TemplateCardEmphasisContent{
			Title: strRef("100"),
			Desc:  strRef("数据含义"),
		},
		QuoteArea: &TemplateCardQuoteArea{
			Type: TemplateCardQuoteAreaTypeUrl,
			Url:  strRef("https://work.weixin.qq.com/?from=openApi"),
		},
		SubTitleText: strRef("下载企业微信还能抢红包！"),
		HorizontalContentList: &[]TemplateCardHorizontalContent{
			{
				KeyName: "邀请人",
				Value:   strRef("张三"),
			},
			{
				KeyName: "企微官网",
				Value:   strRef("点击访问"),
				Type:    TemplateCardHorizontalContentTypeUrl,
				Url:     strRef("https://work.weixin.qq.com/?from=openApi"),
			},
		},
		JumpList: &[]TemplateCardJump{
			{
				Type:  TemplateCardJumpTypeUrl,
				Url:   strRef("https://work.weixin.qq.com/?from=openAp"),
				Title: "企业微信官网",
			},
		},
		CardAction: TemplateCardAction{
			Type: TemplateCardActionTypeUrl,
			Url:  strRef("https://baidu.com"),
		},
	})
	assert.Nil(t, err)
}
