package wxworkbot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUploadMediaAndSendTemplateCardMessage(t *testing.T) {
	botKey := GetTestBotKey()
	if len(botKey) == 0 {
		return
	}
	bot := New(botKey)
	fileBytes := []byte("Here is a string....")
	media, err := bot.UploadMedia(
		"test.txt",
		&fileBytes,
	)
	assert.Nil(t, err)
	assert.NotNil(t, media)

	err = bot.Send(TemplateCard{
		CardType: TemplateCardTypeText,
		Source: &TemplateCardSource{
			IconUrl:   strRef("https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0"),
			Desc:      strRef("企业微信"),
			DescColor: TemplateCardSourceDescColorRed,
		},
		MainTitle: TemplateCardMainTitle{
			Title: strRef("下载 test.txt"),
		},
		HorizontalContentList: &[]TemplateCardHorizontalContent{
			{
				KeyName: "test.txt",
				Value:   strRef("点击下载"),
				Type:    TemplateCardHorizontalContentTypeAttachment,
				MediaID: &media.MediaID,
			},
		},
		CardAction: TemplateCardAction{
			Type: TemplateCardActionTypeUrl,
			Url:  strRef("https://baidu.com"),
		},
	})
	assert.Nil(t, err)
}
