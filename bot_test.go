package wxworkbot

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMarshalText(t *testing.T) {
	text := Text{
		Content:             "广州今日天气：29度，大部分多云，降雨概率：60%",
		MentionedList:       []string{"wangqing", "@all"},
		MentionedMobileList: []string{"13800001111", "@all"},
	}
	msgBytes, err := marshalMessage(text)
	assert.Nil(t, err)
	expected := `{"msgtype":"text","text":{"content":"广州今日天气：29度，大部分多云，降雨概率：60%","mentioned_list":["wangqing","@all"],"mentioned_mobile_list":["13800001111","@all"]}}`
	msg := strings.TrimSuffix(string(msgBytes), "\n")
	assert.Equal(t, expected, msg)
}

func TestMarshalTextMessage(t *testing.T) {
	textMsg := textMessage{
		Text: Text{
			Content:             "广州今日天气：29度，大部分多云，降雨概率：60%",
			MentionedList:       []string{"wangqing", "@all"},
			MentionedMobileList: []string{"13800001111", "@all"},
		},
	}
	msgBytes, err := marshalMessage(textMsg)
	assert.Nil(t, err)
	expected := `{"msgtype":"text","text":{"content":"广州今日天气：29度，大部分多云，降雨概率：60%","mentioned_list":["wangqing","@all"],"mentioned_mobile_list":["13800001111","@all"]}}`
	msg := strings.TrimSuffix(string(msgBytes), "\n")
	assert.Equal(t, expected, msg)
}

func TestMarshalMarkdown(t *testing.T) {
	markdown := Markdown{
		Content: "<font color=\"warning\">233</font>",
	}
	msgBytes, err := marshalMessage(markdown)
	assert.Nil(t, err)
	expected := `{"msgtype":"markdown","markdown":{"content":"<font color=\"warning\">233</font>"}}`
	msg := strings.TrimSuffix(string(msgBytes), "\n")
	assert.Equal(t, expected, msg)
}

func TestMarshalMarkdownMessage(t *testing.T) {
	markdownMsg := markdownMessage{
		Markdown: Markdown{
			Content: "<font color=\"warning\">233</font>",
		},
	}
	msgBytes, err := marshalMessage(markdownMsg)
	assert.Nil(t, err)
	expected := `{"msgtype":"markdown","markdown":{"content":"<font color=\"warning\">233</font>"}}`
	msg := strings.TrimSuffix(string(msgBytes), "\n")
	assert.Equal(t, expected, msg)
}

func TestMarshalImage(t *testing.T) {
	image := Image{
		Base64: "DATA",
		MD5:    "MD5",
	}
	msgBytes, err := marshalMessage(image)
	assert.Nil(t, err)
	expected := `{"msgtype":"image","image":{"base64":"DATA","md5":"MD5"}}`
	msg := strings.TrimSuffix(string(msgBytes), "\n")
	assert.Equal(t, expected, msg)
}

func TestMarshalImageMessage(t *testing.T) {
	imageMsg := imageMessage{
		Image: Image{
			Base64: "DATA",
			MD5:    "MD5",
		},
	}
	msgBytes, err := marshalMessage(imageMsg)
	assert.Nil(t, err)
	expected := `{"msgtype":"image","image":{"base64":"DATA","md5":"MD5"}}`
	msg := strings.TrimSuffix(string(msgBytes), "\n")
	assert.Equal(t, expected, msg)
}

func TestMarshalNews(t *testing.T) {
	news := News{
		Articles: []NewsArticle{
			{
				Title:       "中秋节礼品领取",
				Description: "今年中秋节公司有豪礼相送",
				URL:         "URL",
				PicURL:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
			},
		},
	}
	msgBytes, err := marshalMessage(news)
	assert.Nil(t, err)
	expected := `{"msgtype":"news","news":{"articles":[{"title":"中秋节礼品领取","description":"今年中秋节公司有豪礼相送","url":"URL","picurl":"http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"}]}}`
	msg := strings.TrimSuffix(string(msgBytes), "\n")
	assert.Equal(t, expected, msg)
}

func TestMarshalNewsMessage(t *testing.T) {
	newsMsg := newsMessage{
		News: News{
			Articles: []NewsArticle{
				{
					Title:       "中秋节礼品领取",
					Description: "今年中秋节公司有豪礼相送",
					URL:         "URL",
					PicURL:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
				},
			}},
	}
	msgBytes, err := marshalMessage(newsMsg)
	assert.Nil(t, err)
	expected := `{"msgtype":"news","news":{"articles":[{"title":"中秋节礼品领取","description":"今年中秋节公司有豪礼相送","url":"URL","picurl":"http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"}]}}`
	msg := strings.TrimSuffix(string(msgBytes), "\n")
	assert.Equal(t, expected, msg)
}

func TestMarshalUnsupportedMessage(t *testing.T) {
	text := struct {
		Foo string
	}{
		Foo: "bar",
	}
	_, err := marshalMessage(text)
	assert.Equal(t, ErrUnsupportedMessage, err)
}

func TestSendWithInvalidBotKey(t *testing.T) {
	textMsg := textMessage{
		Text: Text{
			Content:             "广州今日天气：29度，大部分多云，降雨概率：60%",
			MentionedList:       []string{"wangqing", "@all"},
			MentionedMobileList: []string{"13800001111", "@all"},
		},
	}
	bot := New("633a31f6-7f9c-4bc4-97a0-0ec1eefa589")
	err := bot.Send(textMsg)
	assert.NotNil(t, err)
}

//func TestSend(t *testing.T) {
//	textMsg := textMessage{
//		Text:    Text{
//			Content: "测试",
//		},
//	}
//	botKey := os.Getenv("BOTKEY")
//	bot := New(botKey)
//	err := bot.Send(textMsg)
//	bot.Send(textMsg)
//	bot.Send(textMsg)
//	bot.Send(textMsg)
//	assert.Nil(t, err)
//}
