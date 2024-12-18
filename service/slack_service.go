package service

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/adelapazborrero/slack_jack/model"
)

const (
	slackApi = "https://slack.com/api"

	validateEndpoint    = "/auth.test"
	channelListEndpoint = "/conversations.list"
	sendMessageEndpoint = "/chat.postMessage"
	permalinkEndpoint   = "/chat.getPermalink"
)

type SlackService struct {
	apiUrl   string
	SlackBot *model.SlackBot
	Channels *model.ChannelList
}

func NewSlackService(bot *model.SlackBot) *SlackService {
	return &SlackService{
		apiUrl:   slackApi,
		SlackBot: bot,
		Channels: nil,
	}
}

func (serv *SlackService) ValidateBot() error {
	req, err := http.NewRequest(http.MethodPost, slackApi+validateEndpoint, nil)
	if err != nil {
		return errors.New("could not create HTTP request for slack API")
	}

	req.Header.Set("Authorization", "Bearer "+serv.SlackBot.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return errors.New("could not call slack API to validate slack bot token")
	}
	defer response.Body.Close()

	var tokenInfo model.TokenInformation
	err = json.NewDecoder(response.Body).Decode(&tokenInfo)
	if err != nil {
		return errors.New("could not decode JSON response from token validate")
	}

	if !tokenInfo.Ok {
		return errors.New("token was not valid")
	}

	serv.SlackBot.Info = &tokenInfo

	return nil
}

func (serv *SlackService) GetConversationList() error {
	req, err := http.NewRequest(http.MethodPost, slackApi+channelListEndpoint, nil)
	if err != nil {
		return errors.New("could not create HTTP request for slack API")
	}

	req.Header.Set("Authorization", "Bearer "+serv.SlackBot.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return errors.New("could not call slack API to to get conversation list")
	}
	defer response.Body.Close()

	var channelList model.ChannelList
	err = json.NewDecoder(response.Body).Decode(&channelList)
	if err != nil {
		return errors.New("could not decode JSON response for token channel list")
	}

	if !channelList.Ok {
		return errors.New("token does not have permissions to read conversation list")
	}

	serv.Channels = &channelList

	return nil
}
