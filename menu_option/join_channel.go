package menuoption

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adelapazborrero/slack_jack/service"
)

func JoinChannel(slackService *service.SlackService) {
	canReadChannels := true

	if slackService.Channels == nil || len(slackService.Channels.Channels) == 0 {
		fmt.Println("No channels available. Fetching channel list first...")
		err := slackService.GetConversationList()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Select a channel ID manually")
			canReadChannels = false
		}
	}

	if canReadChannels {
		fmt.Println("Select a channel from the list:")
		for _, channel := range slackService.Channels.Channels {
			fmt.Printf("ID: %s, Name: %s\n", channel.ID, channel.Name)
		}
	}

	fmt.Print("Enter Channel ID to join: ")
	reader := bufio.NewReader(os.Stdin)
	channelID, _ := reader.ReadString('\n')
	channelID = strings.TrimSpace(channelID)

	if channelID == "" {
		fmt.Println("Channel ID cannot be empty.")
		return
	}

	err := slackService.JoinChannel(channelID)
	if err != nil {
		log.Println("Failed to join channel:", err)
		return
	}

	fmt.Println("Successfully joined the channel!")
}
