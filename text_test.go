package wxworkbot

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUnmarshalTextMessage(t *testing.T) {
	jsonString := `
		{
			"msgtype": "text",
			"text": {
				"content": "广州今日天气：29度，大部分多云，降雨概率：60%",
				"mentioned_list":["wangqing","@all"],
				"mentioned_mobile_list":["13800001111","@all"]
			}
		}`
	var textMsg textMessage
	err := json.Unmarshal([]byte(jsonString), &textMsg)
	assert.Nil(t, err)
	assert.Equal(t, textMsg.MsgType, "text")
	assert.Equal(t, textMsg.Text.Content, "广州今日天气：29度，大部分多云，降雨概率：60%")
	assert.Equal(t, textMsg.Text.MentionedList, []string{"wangqing", "@all"})
	assert.Equal(t, textMsg.Text.MentionedMobileList, []string{"13800001111", "@all"})
}

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

func TestSendTextMessage(t *testing.T) {
	botKey := GetTestBotKey()
	if len(botKey) == 0 {
		return
	}
	bot := New(botKey)
	err := bot.Send(Text{
		Content:             "测试发送文本消息",
		MentionedList:       []string{"wangqing", "@all"},
		MentionedMobileList: []string{"13800001111", "@all"},
	})
	assert.Nil(t, err)
}
