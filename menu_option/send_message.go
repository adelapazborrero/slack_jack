package menuoption

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adelapazborrero/slack_jack/service"
	"github.com/adelapazborrero/slack_jack/util"
)

func SendMessage(slackService *service.SlackService) {
	channelID := ""
	if slackService.Channels == nil || len(slackService.Channels.Channels) == 0 {
		fmt.Println("No channels available. Fetching channel list first...")
		err := slackService.GetConversationList()
		if err != nil {
			log.Println(err)
			return
		}
	}

	fmt.Println("Select a channel from the list:")
	util.PrintChannelList(slackService.Channels)
	fmt.Print("Enter Channel ID: ")
	reader := bufio.NewReader(os.Stdin)
	channelID, _ = reader.ReadString('\n')
	channelID = strings.TrimSpace(channelID)

	if channelID == "" {
		fmt.Println("Channel ID cannot be empty.")
		return
	}

	fmt.Print("Enter the message to send: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	if message == "" {
		fmt.Println("Message cannot be empty.")
		return
	}

	err := slackService.SendMessage(channelID, message)
	if err != nil {
		log.Println("Failed to send message:", err)
		return
	}

	fmt.Println("Message sent successfully!")
}
