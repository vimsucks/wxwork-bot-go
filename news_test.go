package wxworkbot

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUnmarshalNewsMessage(t *testing.T) {
	jsonString := `
		{
			"msgtype": "news",
			"news": {
			   "articles" : [
				   {
					   "title" : "中秋节礼品领取",
					   "description" : "今年中秋节公司有豪礼相送",
					   "url" : "URL",
					   "picurl" : "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"
				   }
				]
			}
		}
	`
	var newsMsg newsMessage
	err := json.Unmarshal([]byte(jsonString), &newsMsg)
	assert.Nil(t, err)
	assert.Equal(t, newsMsg.MsgType, "news")
	assert.NotEmpty(t, newsMsg.News.Articles)
	article := newsMsg.News.Articles[0]
	assert.Equal(t, article.Title, "中秋节礼品领取")
	assert.Equal(t, article.Description, "今年中秋节公司有豪礼相送")
	assert.Equal(t, article.URL, "URL")
	assert.Equal(t, article.PicURL, "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png")
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
