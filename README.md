<img src="assets/logo.png"/>

# Slack Jack - Slack Bot Token Abuse Tool

Slack Jack is a tool designed for penetration testing purposes. It allows you to interact with Slack channels, send messages, retrieve channel lists, and save sent messages as JSON files. The tool is intended to help ethical hackers explore and test Slack bot token vulnerabilities during security assessments.

This tool provides a command-line interface (CLI) where users can perform various actions related to Slack bot tokens, including interacting with Slack's API endpoints for sending and managing messages.

### Features:

- **Get Channel List**: Fetch and display the list of available Slack channels.
- **Send Message to Channel**: Send messages to selected channels using the Slack bot token.
- **Send Predefined payloads to Channel**: Send built in payloads to selected channels using the Slack bot token.
- **Print Sent Messages**: View the list of messages that have been sent by the bot.
- **Save Sent Messages**: Save the sent messages to a JSON file, with the filename based on the bot user's name and the current date.

## Disclaimer

This tool is a **Work In Progress (WIP)** and is intended **only for educational purposes** and **legal penetration tests**. It should not be used for any unauthorized or malicious activity. Always ensure that you have explicit permission from the target organization before performing any security testing.

By using this tool, you acknowledge and agree to abide by all applicable laws and ethical guidelines related to penetration testing.

## Setup Instructions

### Prerequisites

- Go 1.18+ installed.
- A Slack bot token with appropriate permissions.
- A Slack workspace to test against.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/adelapazborrero/slack_jack.git
   cd slack_jack
   ```

2. Install required Go dependencies:

   ```bash
   go mod tidy
   ```

3. Build the tool:

   ```bash
   go build -o slackjack
   ```

4. Run the tool:

   ```bash
   ./slackjack -t <xoxb-slack-bot-token>
   ```

## Usage

1. **Get Channel List**: Select option 1 to retrieve and display the list of Slack channels available to the bot.
2. **Send Message to Channel**: Select option 2 to send a message to a chosen Slack channel. You'll be prompted to select a channel and input a message.
3. **Send Payload to Channel**: Select option 2 to send a message to a chosen Slack channel. You'll be prompted to select a channel and input an option for predefined payloads, remember to upload the blocks as you will for links etc.
4. **Print Sent Messages**: Select option 3 to print the list of messages that have been sent by the bot.
5. **Save Sent Messages**: Select option 4 to save the sent messages as a JSON file. The filename will include the bot's username and the current date.

# Setting up a test bot

- Create a workspace in slack with any email you want to
- Go to the docs of what Slack API bot tokens are https://api.slack.com/tutorials/tracks/getting-a-token
- Create an app, check the manifest so you can do all the changes in json
- Click on install app and Accept conditions
- Receive and copy your token

- For developing blocks you can use: https://api.slack.com/reference/block-kit/blocks

## Roadmap

- Add option to load payloads from local files (json)
- Add default payloads ready to use (https://api.slack.com/reference/block-kit/block-elements#interactive-components)
- Add enumeration options such as List Users and Chat History

## License

This tool is licensed for educational use and legal penetration tests only. Unauthorized usage is prohibited.
