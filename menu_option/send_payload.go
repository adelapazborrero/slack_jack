package menuoption

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/adelapazborrero/slack_jack/service"
	"github.com/adelapazborrero/slack_jack/util"
)

func SendPredefinedPayload(slackService *service.SlackService) {
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
	util.PrintChannelList(slackService.Channels)
	fmt.Print("Enter Channel ID: ")
	reader := bufio.NewReader(os.Stdin)
	channelID, _ = reader.ReadString('\n')
	channelID = strings.TrimSpace(channelID)

	if channelID == "" {
		fmt.Println("Channel ID cannot be empty.")
		return
	}

	fmt.Println("Available payloads in the 'payloads' folder:")
	files, err := ioutil.ReadDir("./payloads")
	if err != nil {
		log.Println("Failed to read payloads folder:", err)
		return
	}

	var payloadFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			payloadFiles = append(payloadFiles, file.Name())
		}
	}

	if len(payloadFiles) == 0 {
		fmt.Println("No JSON payloads found in the 'payloads' folder.")
		return
	}

	for i, file := range payloadFiles {
		fmt.Printf("%d: %s\n", i+1, file)
	}

	fmt.Print("Select a payload file by number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(payloadFiles) {
		fmt.Println("Invalid selection.")
		return
	}

	selectedFile := filepath.Join("./payloads", payloadFiles[index-1])
	payloadData, err := ioutil.ReadFile(selectedFile)
	if err != nil {
		log.Println("Failed to read the selected payload file:", err)
		return
	}

	var blocks json.RawMessage
	if err := json.Unmarshal(payloadData, &blocks); err != nil {
		fmt.Println("Invalid JSON format:", err)
		return
	}

	err = slackService.SendMessageWithBlocks(channelID, blocks)
	if err != nil {
		log.Println("Failed to send message:", err)
		return
	}

	fmt.Println("Payload sent successfully!")
}
