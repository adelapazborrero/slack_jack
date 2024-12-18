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
				err := slackService.GetConversationList()
				if err != nil {
					log.Println(err)
					return
				}
				PrintChannelList(slackService.Channels)
			},
		},
		// more options here
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
