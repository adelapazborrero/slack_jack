package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adelapazborrero/slack_jack/service"
)

type Menu struct {
	Items map[string]MenuItem
}

type MenuItem struct {
	Description    string
	FunctionToCall func()
}

func NewMenu(items map[string]MenuItem) *Menu {
	return &Menu{Items: items}
}

func BuildMenu(slackService *service.SlackService) *Menu {
	items := map[string]MenuItem{
		"1": {
			Description: "Get Channel List",
			FunctionToCall: func() {
				if slackService.Channels == nil || len(slackService.Channels.Channels) == 0 {
					fmt.Println("Channel list is empty. Fetching channels...")
					err := slackService.GetConversationList()
					if err != nil {
						log.Println(err)
						return
					}
				}
				PrintChannelList(slackService.Channels)
			},
		},
		"2": {
			Description: "Send Message to Channel",
			FunctionToCall: func() {
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
				PrintChannelList(slackService.Channels)
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
			},
		},
		"3": {
			Description: "Print Sent Messages",
			FunctionToCall: func() {
				slackService.PrintSentMessages()
			},
		},
	}

	return NewMenu(items)
}

func (m *Menu) Show() {
	for {
		fmt.Println("\nMenu:")
		for index, item := range m.Items {
			fmt.Printf("%s: %s\n", index, item.Description)
		}
		fmt.Println("q: Quit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nEnter your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "q" {
			fmt.Println("Exiting the tool.")
			return
		}

		if item, exists := m.Items[input]; exists {
			item.FunctionToCall()
		} else {
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
