# wxwork-bot-go

企业微信群机器人接口 Golang 封装

## Quick Start
```go
package main

import (
	"fmt"
    "github.com/vimsucks/wxwork-bot-go"
)

func main() {
    bot := wxworkbot.New("YOUR_BOT_HERE")
    // or Markdown, Image, News
    text := wxworkbot.Text{
    	Content: "Hello World",
    	MentionedList: []string{"foo", "bar"},
    	MentionedMobileList: []string{"@all"},
    }
    err := bot.Send(text)
    if err != nil {
    	log.Fatal(err)
    }
    fmt.Println("success!")
}
```