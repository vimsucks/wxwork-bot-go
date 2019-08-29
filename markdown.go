package wxworkbot

type markdownMessage struct {
	message
	Markdown Markdown `json:"markdown"`
}

type Markdown struct {
	Content string `json:"content"`
}
