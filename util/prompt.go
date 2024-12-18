package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Prompt struct {
}

func NewPrompt() *Prompt {
	return &Prompt{}
}

func (p Prompt) EnumerateChannels() bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Do you want to enumerate the channels? (Y/n): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "" || input == "y" {
			return true
		} else if input == "n" {
			return false
		} else {
			fmt.Println("Invalid input. Please enter 'Y' or 'n'.")
		}
	}
}
