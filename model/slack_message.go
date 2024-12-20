package model

type SlackMessageResponse struct {
	Ok      bool         `json:"ok"`
	Channel string       `json:"channel"`
	Ts      string       `json:"ts"`
	Message SlackMessage `json:"message"`
	Blocks  []SlackBlock `json:"blocks"`
}

type SlackMessagesResponse struct {
	Ok                  bool           `json:"ok"`
	Messages            []SlackMessage `json:"messages"`
	HasMore             bool           `json:"has_more"`
	PinCount            int            `json:"pin_count"`
	Warning             string         `json:"warning"`
	ChannelActionsTs    *string        `json:"channel_actions_ts"`
	ChannelActionsCount int            `json:"channel_actions_count"`
}

type SlackMessageMap struct {
	Messages map[string][]SlackSentMessage `json:"messages"`
}

type SlackPermalinkResponse struct {
	Ok        bool   `json:"ok"`
	Permalink string `json:"permalink"`
	Channel   string `json:"channel"`
}

type SlackSentMessage struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Ts        string `json:"ts"`
	Permalink string `json:"permalink"`
}

type SlackMessage struct {
	User       string          `json:"user"`
	Type       string          `json:"type"`
	Subtype    string          `json:"subtype"`
	Ts         string          `json:"ts"`
	BotID      string          `json:"bot_id"`
	AppID      string          `json:"app_id"`
	Text       string          `json:"text"`
	Team       string          `json:"team"`
	BotProfile SlackBotProfile `json:"bot_profile"`
	Blocks     []SlackBlock    `json:"blocks"`
}

type SlackBotProfile struct {
	ID      string     `json:"id"`
	AppID   string     `json:"app_id"`
	Name    string     `json:"name"`
	Icons   SlackIcons `json:"icons"`
	Deleted bool       `json:"deleted"`
	Updated int        `json:"updated"`
	TeamID  string     `json:"team_id"`
}

type SlackIcons struct {
	Image36 string `json:"image_36"`
	Image48 string `json:"image_48"`
	Image72 string `json:"image_72"`
}

type SlackBlock struct {
	Type     string         `json:"type"`
	BlockID  string         `json:"block_id"`
	Elements []SlackElement `json:"elements"`
}

type SlackElement struct {
	Type     string        `json:"type"`
	Elements []TextElement `json:"elements"`
}

type TextElement struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
