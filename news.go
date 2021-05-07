package sdk

type newsMessage struct {
	messagetype
	News News `json:"news"`
}

type News struct {
	Articles []NewsArticle `json:"articles"`
}

type NewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}


