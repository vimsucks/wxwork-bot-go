package wxworkbot

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUnmarshalMarkdownMessage(t *testing.T) {
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
