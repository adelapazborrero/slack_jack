package menuoption

import "github.com/adelapazborrero/slack_jack/service"

func PrintSentMessages(slackService *service.SlackService) {
	slackService.PrintSentMessages()
}
