package wxworkbot

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextMessage(t *testing.T) {
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
