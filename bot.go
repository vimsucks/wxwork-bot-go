package wxworkbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func init() {
	client = &http.Client{
		Timeout: 5*time.Second,
	}
}

const (
	defaultWebHookUrlTemplate = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"
)

var (
	ErrUnsupportedMessage = errors.New("尚不支持的消息类型")
	client *http.Client
)

type WxWorkBot struct {
	Key        string
	WebHookUrl string
}

// 创建一个新的机器人实例
func New(botKey string) *WxWorkBot {
	bot := WxWorkBot{
		Key: botKey,
		// 直接拼接出接口 URL
		WebHookUrl: fmt.Sprintf(defaultWebHookUrlTemplate, botKey),
	}
	return &bot
}

// 发送消息，允许的参数类型：Text、Markdown、Image、News
func (bot *WxWorkBot) Send(msg interface{}) error {
	msgBytes, err := marshalMessage(msg)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, bot.WebHookUrl, bytes.NewBuffer(msgBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var wxWorkResp wxWorkResponse
	err = json.Unmarshal(body, &wxWorkResp)
	if err != nil {
		return err
	}
	if wxWorkResp.ErrorCode != 0 && wxWorkResp.ErrorMessage != "" {
		return errors.New(string(body))
	}
	return nil
}

// 防止 < > 等 HTML 字符在 json.marshal 时被 escape
func marshal(msg interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", "")
	err := jsonEncoder.Encode(msg)
	if err != nil {
		return nil, nil
	}
	return buf.Bytes(), nil
}

// 将消息包装成企信接口要求的格式，返回 json bytes
func marshalMessage(msg interface{}) ([]byte, error) {
	if text, ok := msg.(Text); ok {
		textMsg := textMessage{
			message: message{MsgType: "text"},
			Text:    text,
		}
		return marshal(textMsg)
	}
	if textMsg, ok := msg.(textMessage); ok {
		textMsg.MsgType = "text"
		return marshal(textMsg)
	}
	if markdown, ok := msg.(Markdown); ok {
		markdownMsg := markdownMessage{
			message:  message{MsgType: "markdown"},
			Markdown: markdown,
		}
		return marshal(markdownMsg)
	}
	if markdownMsg, ok := msg.(markdownMessage); ok {
		markdownMsg.MsgType = "markdown"
		return marshal(markdownMsg)
	}
	if image, ok := msg.(Image); ok {
		imageMsg := imageMessage{
			message: message{MsgType: "image"},
			Image:   image,
		}
		return marshal(imageMsg)
	}
	if imageMsg, ok := msg.(imageMessage); ok {
		imageMsg.MsgType = "image"
		return marshal(imageMsg)
	}
	if news, ok := msg.(News); ok {
		newsMsg := newsMessage{
			message: message{MsgType: "news"},
			News:    news,
		}
		return marshal(newsMsg)
	}
	if newsMsg, ok := msg.(newsMessage); ok {
		newsMsg.MsgType = "news"
		return marshal(newsMsg)
	}
	return nil, ErrUnsupportedMessage
}
