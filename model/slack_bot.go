package model

import (
	"errors"
	"strings"
)

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

type SlackBot struct {
	Token string
	Info  *TokenInformation
}

func NewSlackBot(token string) *SlackBot {
	return &SlackBot{
		Token: token,
		Info:  nil,
	}
}

func (s SlackBot) Validate() error {

	if !strings.Contains(s.Token, "xoxb") && !strings.Contains(s.Token, "xoxp") {
		return errors.New("token must start with xoxb or xoxp")
	}

	return nil
}
