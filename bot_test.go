package wxworkbot

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
	"time"
)

func GetTestBotKey() string {
	return os.Getenv("WXWORK_BOT_KEY")
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
	bot := New("这是一个错误的 BOT KEY")
	err := bot.Send(textMsg)
	assert.NotNil(t, err)
}

func TestWithCustomHttpClient(t *testing.T) {
	botKey := GetTestBotKey()
	if len(botKey) == 0 {
		return
	}
	bot := WxWorkBot{
		Key: botKey,
		Client: &http.Client{
			Timeout: 1 * time.Second,
		},
	}
	err := bot.Send(Text{
		Content:             "广州今日天气：29度，大部分多云，降雨概率：60%",
		MentionedList:       []string{"wangqing", "@all"},
		MentionedMobileList: []string{"13800001111", "@all"},
	})
	assert.Nil(t, err)
}
