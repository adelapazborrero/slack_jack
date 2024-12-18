package main

import (
	"log"

	"github.com/adelapazborrero/slack_jack/model"
	"github.com/adelapazborrero/slack_jack/service"
	"github.com/adelapazborrero/slack_jack/util"
)

func main() {
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

	util.PrintTokenInformation(slackService.SlackBot.Info)

	prompt := util.NewPrompt()

	enumerateChannels := prompt.EnumerateChannels()

	if enumerateChannels {
		err = slackService.GetConversationList()
		if err != nil {
			log.Fatal(err)
		}
		util.PrintChannelList(slackService.Channels)
	}

}
