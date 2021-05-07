package sdk

type markdownMessage struct {
	messagetype
	Markdown Markdown `json:"markdown"`
}

type Markdown struct {
	Content string `json:"content"`
}