package menuoption

import (
	"fmt"
	"log"

	"github.com/adelapazborrero/slack_jack/service"
	"github.com/adelapazborrero/slack_jack/util"
)

func GetChannelList(slackService *service.SlackService) {

	if slackService.Channels == nil || len(slackService.Channels.Channels) == 0 {
		fmt.Println("Channel list is empty. Fetching channels...")
		err := slackService.GetConversationList()
		if err != nil {
			log.Println(err)
			return
		}
	}
	util.PrintChannelList(slackService.Channels)
}
