package wxworkbot

type MarkdownMessage struct {
	Message
	Markdown Markdown `json:"markdown"`
}

type Markdown struct {
	Content string `json:"content"`
}
