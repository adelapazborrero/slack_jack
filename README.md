<img src="assets/logo.png"/>

# Slack Jack - Slack Bot Token Abuse

Slack Jack is a penetration testing tool designed for ethical hacking and security testing purposes. It allows you to hijack a Slack bot using its token (e.g., xoxb or xoxp) and perform various enumeration and exploitation activities, depending on the bot's assigned permissions.

Convincing a user to interact with a malicious link or payload can often be the most challenging part of gaining initial access. By impersonating a trusted bot, Slack Jack opens up numerous possibilities for social engineering attacks. For example, combining this tool with Evilginx could be an effective way to capture credentials.

This tool provides a command-line interface (CLI) that enables users to interact with Slack's API endpoints, facilitating actions like sending messages, managing bot activities, and more.

### Features:

- **Get Channel List**: Retrieve and display the list of Slack channels accessible to the bot.
- **Send Message to Channel**: Use the Slack bot token to send messages to specified channels.
- **Send Predefined payloads to Channel**: Deploy built-in payloads to target channels using the Slack bot token.
- **Print Sent Messages**: Display a history of messages sent by the bot.
- **Save Sent Messages**: Export sent messages to a JSON file, with the filename based on the bot user's name and the current date.
- **Join Channel**: Join a channel using its Channel ID, provided the bot has the necessary permissions.
- **Print Chat History**: Extract and display a specified number of messages from a channel, if the bot has permission to access the chat history.

### Demo against test target

Here's an example of Slack Jack in action against a test target:

<img src="assets/demo.gif"/>

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

After initializing the tool, it will validate the provided bot token. If the token is valid, you will be presented with a menu of available options. Select your desired actions and follow the prompts. Note that some commands may not work if the bot lacks the necessary permissions.

# Setting up a test bot

- Create a Slack workspace using any email address.
- Review the Slack API bot token documentation to understand how to generate a token. https://api.slack.com/tutorials/tracks/getting-a-token
- Create an app in Slack and use the manifest to make changes in JSON format.
- Install the app and accept the terms and conditions.
- Copy the generated bot token.
- For developing and testing interactive blocks, refer to: https://api.slack.com/reference/block-kit/blocks

## Roadmap

- Add option to load payloads from local files (json)
- Add default payloads ready to use (https://api.slack.com/reference/block-kit/block-elements#interactive-components)
- Add enumeration options such as List Users

## License

This tool is licensed for educational use and legal penetration tests only. Unauthorized usage is prohibited.
