package menuoption

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/adelapazborrero/slack_jack/service"
)

func GetChannelHistory(slackService *service.SlackService) {
	if slackService.Channels == nil || len(slackService.Channels.Channels) == 0 {
		fmt.Println("No channels available. Fetching channel list first...")
		err := slackService.GetConversationList()
		if err != nil {
			log.Println(err)
			return
		}
	}

	fmt.Println("Select a channel from the list:")
	for _, channel := range slackService.Channels.Channels {
		fmt.Printf("ID: %s, Name: %s\n", channel.ID, channel.Name)
	}

	fmt.Print("Enter Channel ID to fetch history: ")
	reader := bufio.NewReader(os.Stdin)
	channelID, _ := reader.ReadString('\n')
	channelID = strings.TrimSpace(channelID)

	if channelID == "" {
		fmt.Println("Channel ID cannot be empty.")
		return
	}

	fmt.Print("Enter the number of messages to retrieve: ")
	limitInput, _ := reader.ReadString('\n')
	limitInput = strings.TrimSpace(limitInput)
	limit, err := strconv.Atoi(limitInput)
	if err != nil || limit <= 0 {
		fmt.Println("Invalid limit. Please enter a positive number.")
		return
	}

	messages, err := slackService.GetChannelHistory(channelID, limit)
	if err != nil {
		log.Println("Failed to fetch channel history:", err)
		return
	}

	slices.Reverse(messages)

	fmt.Println("Channel History:")
	for _, msg := range messages {
		fmt.Printf("%s: %s\n", msg.User, msg.Text)
	}
}
