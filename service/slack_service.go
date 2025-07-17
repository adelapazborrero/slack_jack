package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/adelapazborrero/slack_jack/model"
)

const (
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

func NewSlackService(bot *model.SlackBot, apiUrl string) *SlackService {
	return &SlackService{
		apiUrl:   apiUrl,
		SlackBot: bot,
		Channels: &model.ChannelList{
			Channels: make([]model.Channel, 0),
		},
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
			fmt.Printf("  Message #%d:\n", i+1)
			fmt.Printf("    Text: %s\n", msg.Text)
			fmt.Printf("    Timestamp: %s\n", msg.Ts)
			fmt.Printf("    Permalink: %s\n", msg.Permalink)
		}
	}
}

func (serv *SlackService) ValidateBot() error {
	req, err := http.NewRequest(http.MethodPost, serv.apiUrl+validateEndpoint, nil)
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
	req, err := http.NewRequest(http.MethodPost, serv.apiUrl+channelListEndpoint, nil)
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

	req, err := http.NewRequest(http.MethodPost, serv.apiUrl+sendMessageEndpoint, bytes.NewBuffer(jsonPayload))
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

	messageTs := slackResponse.Message.Ts

	permalinkResponse, err := serv.getPermalink(channelID, messageTs)
	if err != nil {
		return err
	}

	sentMessage := model.SlackSentMessage{
		ID:        slackResponse.Message.Ts,
		Text:      slackResponse.Message.Text,
		Ts:        slackResponse.Message.Ts,
		Permalink: permalinkResponse.Permalink,
	}

	serv.Messages.Messages[channelID] = append(serv.Messages.Messages[channelID], sentMessage)
	return nil
}

func (serv *SlackService) SendMessageWithBlocks(channelID string, blocks json.RawMessage) error {
	var parsedBlocks []map[string]interface{}
	err := json.Unmarshal(blocks, &parsedBlocks)
	if err != nil {
		return fmt.Errorf("invalid blocks format: %v", err)
	}

	payload := map[string]interface{}{
		"channel": channelID,
		"blocks":  parsedBlocks,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return errors.New("could not marshal message payload to JSON")
	}

	req, err := http.NewRequest(http.MethodPost, serv.apiUrl+sendMessageEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return errors.New("could not create HTTP request for sending message with blocks")
	}

	req.Header.Set("Authorization", "Bearer "+serv.SlackBot.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return errors.New("could not call Slack API to send message with blocks")
	}
	defer response.Body.Close()

	var slackResponse model.SlackMessageResponse
	err = json.NewDecoder(response.Body).Decode(&slackResponse)
	if err != nil {
		return fmt.Errorf("could not decode JSON response for sendMessageWithBlocks: %s", err)
	}

	if !slackResponse.Ok {
		return fmt.Errorf("failed to send message with blocks: %s", slackResponse.Message.Text)
	}

	messageTs := slackResponse.Message.Ts
	permalinkResponse, err := serv.getPermalink(channelID, messageTs)
	if err != nil {
		return err
	}

	sentMessage := model.SlackSentMessage{
		ID:        slackResponse.Message.Ts,
		Text:      "Blocks Sent",
		Ts:        slackResponse.Message.Ts,
		Permalink: permalinkResponse.Permalink,
	}

	serv.Messages.Messages[channelID] = append(serv.Messages.Messages[channelID], sentMessage)
	return nil
}

func (s *SlackService) getPermalink(channelID, messageTs string) (*model.SlackPermalinkResponse, error) {
	formData := url.Values{}
	formData.Set("channel", channelID)
	formData.Set("message_ts", messageTs)

	req, err := http.NewRequest("POST", s.apiUrl+permalinkEndpoint, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create permalink request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+s.SlackBot.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get permalink: %v", err)
	}
	defer resp.Body.Close()

	var permalinkResp model.SlackPermalinkResponse
	if err := json.NewDecoder(resp.Body).Decode(&permalinkResp); err != nil {
		return nil, fmt.Errorf("failed to decode permalink response: %v", err)
	}

	if !permalinkResp.Ok {
		return nil, fmt.Errorf("failed to retrieve permalink: %v", permalinkResp)
	}

	return &permalinkResp, nil
}

func (serv *SlackService) JoinChannel(channelID string) error {
	joinEndpoint := "/conversations.join"

	payload := map[string]string{
		"channel": channelID,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return errors.New("could not marshal join channel payload to JSON")
	}

	req, err := http.NewRequest(http.MethodPost, serv.apiUrl+joinEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return errors.New("could not create HTTP request for joining channel")
	}

	req.Header.Set("Authorization", "Bearer "+serv.SlackBot.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return errors.New("could not call Slack API to join channel")
	}
	defer response.Body.Close()

	var joinResponse struct {
		Ok    bool   `json:"ok"`
		Error string `json:"error,omitempty"`
	}

	err = json.NewDecoder(response.Body).Decode(&joinResponse)
	if err != nil {
		return errors.New("could not decode JSON response for joinChannel")
	}

	if !joinResponse.Ok {
		return fmt.Errorf("failed to join channel: %s", joinResponse.Error)
	}

	return nil
}

func (serv *SlackService) GetChannelHistory(channelID string, limit int) ([]model.SlackMessage, error) {
	historyEndpoint := "/conversations.history"

	params := url.Values{}
	params.Set("channel", channelID)
	params.Set("limit", fmt.Sprintf("%d", limit))

	req, err := http.NewRequest(http.MethodGet, serv.apiUrl+historyEndpoint+"?"+params.Encode(), nil)
	if err != nil {
		return nil, errors.New("could not create HTTP request for fetching channel history")
	}

	req.Header.Set("Authorization", "Bearer "+serv.SlackBot.Token)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, errors.New("could not call Slack API to get channel history")
	}
	defer response.Body.Close()

	var historyResponse struct {
		Ok       bool                 `json:"ok"`
		Messages []model.SlackMessage `json:"messages"`
		Error    string               `json:"error,omitempty"`
	}

	err = json.NewDecoder(response.Body).Decode(&historyResponse)
	if err != nil {
		return nil, errors.New("could not decode JSON response for getChannelHistory")
	}

	if !historyResponse.Ok {
		return nil, fmt.Errorf("failed to fetch channel history: %s", historyResponse.Error)
	}

	return historyResponse.Messages, nil
}
