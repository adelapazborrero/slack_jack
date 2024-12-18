package util

import (
	"fmt"

	"github.com/adelapazborrero/slack_jack/model"
)

func PrintChannelList(chanList *model.ChannelList) {
	if chanList == nil || len(chanList.Channels) == 0 {
		fmt.Println("No channels available.")
		return
	}

	fmt.Println("Channel List:")
	fmt.Println("-------------")
	for _, channel := range chanList.Channels {
		fmt.Printf("ID:          %s\n", channel.ID)
		fmt.Printf("Name:        %s\n", channel.Name)
		fmt.Printf("Private:     %v\n", channel.IsPrivate)
		fmt.Printf("Members:     %d\n", channel.NumMembers)
		if channel.Topic.Value != "" {
			fmt.Printf("Topic:       %s\n", channel.Topic.Value)
		}
		if channel.Purpose.Value != "" {
			fmt.Printf("Purpose:     %s\n", channel.Purpose.Value)
		}
		fmt.Println("-------------")
	}
}
