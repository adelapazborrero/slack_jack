package model

type SlackMessageResponse struct {
	Ok      bool         `json:"ok"`
	Channel string       `json:"channel"`
	Ts      string       `json:"ts"`
	Message SlackMessage `json:"message"`
	Blocks  []SlackBlock `json:"blocks"`
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
	Ts         string          `json:"ts"`
	BotID      string          `json:"bot_id"`
	AppID      string          `json:"app_id"`
	Text       string          `json:"text"`
	Team       string          `json:"team"`
	BotProfile SlackBotProfile `json:"bot_profile"`
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
