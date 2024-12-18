package main

import (
	"fmt"
	"log"

	"github.com/adelapazborrero/slack_jack/model"
	"github.com/adelapazborrero/slack_jack/service"
)

func main() {
	//To be passed as argument
	slackToken := ""

	slackBot := model.NewSlackBot(slackToken)

	err := slackBot.Validate()
	if err != nil {
		log.Fatal(err)
		return
	}

	slackService := service.NewSlackService(slackBot)

	err = slackService.ValidateBot()
	if err != nil {
		log.Fatal(err)
	}

	err = slackService.GetConversationList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(slackService.SlackBot.Info)
	fmt.Println(slackService.Channels)
}
