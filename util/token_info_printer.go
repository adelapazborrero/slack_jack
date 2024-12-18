package util

import (
	"fmt"

	"github.com/adelapazborrero/slack_jack/model"
)

func PrintTokenInformation(information *model.TokenInformation) {
	if information == nil {
		fmt.Println("No token information available.")
		return
	}

	fmt.Println("Token Information:")
	fmt.Println("-------------------")
	fmt.Printf("Team:       %s (ID: %s)\n", information.Team, information.TeamID)
	fmt.Printf("User:       %s (ID: %s)\n", information.User, information.UserID)
	fmt.Printf("Bot ID:     %s\n", information.BotID)
	fmt.Printf("Enterprise: %v\n", information.IsEnterprise)
	fmt.Printf("URL:        %s\n", information.Url)
	fmt.Println("-------------------")
}
