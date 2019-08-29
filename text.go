package wxworkbot

type textMessage struct {
	message
	Text Text `json:"text"`
}

type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}
