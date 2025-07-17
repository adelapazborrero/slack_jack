package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/adelapazborrero/slack_jack/menu"
	"github.com/adelapazborrero/slack_jack/model"
	"github.com/adelapazborrero/slack_jack/service"
	"github.com/adelapazborrero/slack_jack/util"
)

func main() {
	slackToken := flag.String("t", "", "Slack Bot Token")
	flag.StringVar(slackToken, "token", "", "Slack Bot Token")
	slackApiUrl := flag.String("api", "https://slack.com/api", "Slack API base URL")
	flag.Parse()

	if *slackToken == "" {
		fmt.Println("Slack Bot Token is required. Please use the -t or --token flag.")
		os.Exit(1)
	}

	slackBot := model.NewSlackBot(*slackToken)

	err := slackBot.Validate()
	if err != nil {
		log.Fatal(err)
		return
	}

	slackService := service.NewSlackService(slackBot, *slackApiUrl)

	err = slackService.ValidateBot()
	if err != nil {
		log.Fatal(err)
	}

	util.PrintTokenInformation(slackService.SlackBot.Info)

	menu := menu.Build(slackService)
	menu.Show()
}
