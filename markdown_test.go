package wxworkbot

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarkdownMessage(t *testing.T) {
	jsonString := `
		{
			"msgtype": "markdown",
			"markdown": {
				"content": "<font color=\"warning\">233</font>"
			}
		}`
	var markdownMsg markdownMessage
	err := json.Unmarshal([]byte(jsonString), &markdownMsg)
	assert.Nil(t, err)
	assert.Equal(t, markdownMsg.MsgType, "markdown")
	assert.Equal(t, markdownMsg.Markdown.Content,
		"<font color=\"warning\">233</font>")
}
