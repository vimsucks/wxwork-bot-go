package wxworkbot

const (
	// TemplateCardTypeText 文本通知类型的模板卡片
	TemplateCardTypeText = "text_notice"
	// TemplateCardTypeNews 图文展示类型的模板卡片
	TemplateCardTypeNews = "news_notice"
)

type templateCardMessage struct {
	message
	TemplateCard TemplateCard `json:"template_card"`
}

type TemplateCard struct {
	CardType              string                           `json:"card_type"`
	Source                *TemplateCardSource              `json:"source"`
	MainTitle             TemplateCardMainTitle            `json:"main_title"`
	CardImage             *TemplateCardImage               `json:"card_image"`
	ImageTextArea         *TemplateCardImageTextArea       `json:"image_text_area"`
	EmphasisContent       *TemplateCardEmphasisContent     `json:"emphasis_content"`
	QuoteArea             *TemplateCardQuoteArea           `json:"quote_area"`
	SubTitleText          *string                          `json:"sub_title_text"`
	HorizontalContentList *[]TemplateCardHorizontalContent `json:"horizontal_content_list"`
	JumpList              *[]TemplateCardJump              `json:"jump_list"`
	CardAction            TemplateCardAction               `json:"card_action"`
}

type TemplateCardSourceDescColor int

// 来源文字的颜色，目前支持：0(默认) 灰色，1 黑色，2 红色，3 绿色
const (
	TemplateCardSourceDescColorGrey  = 0
	TemplateCardSourceDescColorBlack = 1
	TemplateCardSourceDescColorRed   = 2
	TemplateCardSourceDescColorGreen = 3
)

type TemplateCardSource struct {
	IconUrl   *string `json:"icon_url"`
	Desc      *string `json:"desc"`
	DescColor int     `json:"desc_color"`
}

type TemplateCardMainTitle struct {
	Title *string `json:"title"`
	Desc  *string `json:"desc"`
}

type TemplateCardEmphasisContent struct {
	Title *string `json:"title"`
	Desc  *string `json:"desc"`
}

type TemplateCardQuoteAreaType int

// 引用文献样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
const (
	TemplateCardQuoteAreaTypeNone    TemplateCardQuoteAreaType = 0
	TemplateCardQuoteAreaTypeUrl     TemplateCardQuoteAreaType = 1
	TemplateCardQuoteAreaTypeMiniApp TemplateCardQuoteAreaType = 2
)

type TemplateCardQuoteArea struct {
	Type      TemplateCardQuoteAreaType `json:"type"`
	Url       *string                   `json:"url"`
	AppID     *string                   `json:"appid"`
	PagePath  *string                   `json:"pagepath"`
	Title     *string                   `json:"title"`
	QuoteText *string                   `json:"quote_text"`
}

type TemplateCardHorizontalContentType int

// 链接类型，0或不填代表是普通文本，1 代表跳转url，2 代表下载附件，3 代表@员工
const (
	TemplateCardHorizontalContentTypeText       TemplateCardHorizontalContentType = 0
	TemplateCardHorizontalContentTypeUrl        TemplateCardHorizontalContentType = 1
	TemplateCardHorizontalContentTypeAttachment TemplateCardHorizontalContentType = 2
	TemplateCardHorizontalContentTypeMention    TemplateCardHorizontalContentType = 3
)

type TemplateCardHorizontalContent struct {
	KeyName string                            `json:"keyname"`
	Value   *string                           `json:"value"`
	Type    TemplateCardHorizontalContentType `json:"type"`
	Url     *string                           `json:"url"`
	MediaID *string                           `json:"media_id"`
	UserID  *string                           `json:"userid"`
}

type TemplateCardJumpType int

// 跳转链接类型，0或不填代表不是链接，1 代表跳转url，2 代表跳转小程序
const (
	TemplateCardJumpTypeNone    TemplateCardJumpType = 0
	TemplateCardJumpTypeUrl     TemplateCardJumpType = 1
	TemplateCardJumpTypeMiniApp TemplateCardJumpType = 2
)

type TemplateCardJump struct {
	Type     TemplateCardJumpType `json:"type"`
	Url      *string              `json:"url"`
	Title    string               `json:"title"`
	AppID    *string              `json:"appid"`
	PagePath *string              `json:"pagepath"`
}
type TemplateCardActionType int

// 卡片跳转类型，1 代表跳转url，2 代表打开小程序。text_notice模版卡片中该字段取值范围为[1,2]
const (
	TemplateCardActionTypeUrl     TemplateCardActionType = 1
	TemplateCardActionTypeMiniApp TemplateCardActionType = 2
)

type TemplateCardAction struct {
	Type     TemplateCardActionType `json:"type"`
	Url      *string                `json:"url"`
	AppID    *string                `json:"appid"`
	PagePath *string                `json:"pagepath"`
}

type TemplateCardImage struct {
	Url string `json:"url"`
	// 图片的宽高比，宽高比要小于2.25，大于1.3，不填该参数默认1.3
	AspectRatio *float64 `json:"aspect_ratio"`
}

type TemplateCardImageTextAreaType int

const (
	TemplateCardImageTextAreaTypeNone    TemplateCardImageTextAreaType = 0
	TemplateCardImageTextAreaTypeUrl     TemplateCardImageTextAreaType = 1
	TemplateCardImageTextAreaTypeMiniApp TemplateCardImageTextAreaType = 2
)

type TemplateCardImageTextArea struct {
	Type     TemplateCardImageTextAreaType `json:"type"`
	Url      *string                       `json:"url"`
	AppID    *string                       `json:"appid"`
	PagePath *string                       `json:"pagepath"`
	Title    *string                       `json:"title"`
	Desc     *string                       `json:"desc"`
	ImageUrl string                        `json:"image_url"`
}
