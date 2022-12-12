package wxworkbot

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUnmarshalImageMessage(t *testing.T) {
	jsonString := `
		{
			"msgtype": "image",
			"image": {
				"base64": "DATA",
				"md5": "MD5"
			}
		}
	`
	var imageMsg imageMessage
	err := json.Unmarshal([]byte(jsonString), &imageMsg)
	assert.Nil(t, err)
	assert.Equal(t, imageMsg.MsgType, "image")
	assert.Equal(t, imageMsg.Image.Base64, "DATA")
	assert.Equal(t, imageMsg.Image.MD5, "MD5")
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
