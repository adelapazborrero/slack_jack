package menuoption

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adelapazborrero/slack_jack/service"
)

func SaveSentMessages(slackService *service.SlackService) {
	if len(slackService.Messages.Messages) == 0 {
		fmt.Println("No messages to save.")
		return
	}

	botUserName := slackService.SlackBot.Info.User

	currentDate := time.Now().Format("2006-01-02")

	fileName := fmt.Sprintf("%s_sent_messages_%s.json", botUserName, currentDate)

	file, err := os.Create(fileName)
	if err != nil {
		log.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	data, err := json.MarshalIndent(slackService.Messages, "", "  ")
	if err != nil {
		log.Println("Failed to marshal messages to JSON:", err)
		return
	}

	_, err = file.Write(data)
	if err != nil {
		log.Println("Failed to write data to file:", err)
		return
	}

	fmt.Printf("Sent messages saved to %s\n", fileName)
}
