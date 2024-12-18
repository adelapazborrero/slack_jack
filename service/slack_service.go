package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	Messages *model.SlackMessageMap
}

func NewSlackService(bot *model.SlackBot) *SlackService {
	return &SlackService{
		apiUrl:   slackApi,
		SlackBot: bot,
		Channels: nil,
		Messages: &model.SlackMessageMap{
			Messages: make(map[string][]model.SlackSentMessage),
		},
	}
}

func (serv *SlackService) PrintSentMessages() {
	if len(serv.Messages.Messages) == 0 {
		fmt.Println("No messages sent yet.")
		return
	}

	for channelID, messages := range serv.Messages.Messages {
		fmt.Printf("Channel: %s\n", channelID)
		for i, msg := range messages {
			fmt.Printf("  Message #%d: %s (ID: %s)\n", i+1, msg.Text, msg.ID)
		}
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

func (serv *SlackService) SendMessage(channelID, message string) error {
	payload := map[string]string{
		"channel": channelID,
		"text":    message,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return errors.New("could not marshal message payload to JSON")
	}

	req, err := http.NewRequest(http.MethodPost, slackApi+sendMessageEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return errors.New("could not create HTTP request for sending message")
	}

	req.Header.Set("Authorization", "Bearer "+serv.SlackBot.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return errors.New("could not call slack API to send message")
	}
	defer response.Body.Close()

	var slackResponse model.SlackMessageResponse
	err = json.NewDecoder(response.Body).Decode(&slackResponse)
	if err != nil {
		return errors.New("could not decode JSON response for sendMessage")
	}

	if !slackResponse.Ok {
		return errors.New("failed to send message: " + slackResponse.Message.Text)
	}

	sentMessage := model.SlackSentMessage{
		ID:   slackResponse.Message.Ts,
		Text: slackResponse.Message.Text,
		Ts:   slackResponse.Message.Ts,
	}

	serv.Messages.Messages[channelID] = append(serv.Messages.Messages[channelID], sentMessage)
	return nil
}
