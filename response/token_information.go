package response

type TokenInformation struct {
	Ok           bool   `json:"ok"`
	Url          string `json:"url"`
	Team         string `json:"team"`
	User         string `json:"user"`
	TeamID       string `json:"team_id"`
	UserID       string `json:"user_id"`
	BotID        string `json:"bot_id"`
	IsEnterprise bool   `json:"is_enterprise_install"`
}
