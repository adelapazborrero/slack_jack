package menu

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	menuoption "github.com/adelapazborrero/slack_jack/menu_option"
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

func Build(slackService *service.SlackService) *Menu {
	items := map[string]MenuItem{
		"1": {
			Description: "Get Channel List",
			FunctionToCall: func() {
				menuoption.GetChannelList(slackService)
			},
		},
		"2": {
			Description: "Send Message to Channel",
			FunctionToCall: func() {
				menuoption.SendMessage(slackService)
			},
		},
		"3": {
			Description: "Print Sent Messages",
			FunctionToCall: func() {
				menuoption.PrintSentMessages(slackService)
			},
		},
		"4": {
			Description: "Save Sent Messages as JSON",
			FunctionToCall: func() {
				menuoption.SaveSentMessages(slackService)
			},
		},
	}

	return NewMenu(items)
}

func (m *Menu) Show() {
	for {
		fmt.Println("\nMenu:")

		var keys []string
		for key := range m.Items {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {
			item := m.Items[key]
			fmt.Printf("%s: %s\n", key, item.Description)
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
