package wxworkbot

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageMessage(t *testing.T) {
	jsonString := `
		{
			"msgtype": "image",
			"image": {
				"base64": "DATA",
				"md5": "MD5"
			}
		}
	`
	var imageMsg ImageMessage
	err := json.Unmarshal([]byte(jsonString), &imageMsg)
	assert.Nil(t, err)
	assert.Equal(t, imageMsg.MsgType, "image")
	assert.Equal(t, imageMsg.Image.Base64, "DATA")
	assert.Equal(t, imageMsg.Image.MD5, "MD5")
}
