# Project 1: Slack Bot with Slacker Library
## YOB Calculator Bot
- This Slack bot calculates the age based on the provided year of birth (YOB).
- This project implements a simple Slack bot using the Slacker library in Golang. The bot listens for a specific command (my yob is <year>) and calculates the user's age based on the provided birth year.

## Setup
### 1)Install dependencies:

- go get -u github.com/shomali11/slacker

### 2)Set your Slack Bot Token and App Token as environment variables:

- export SLACK_BOT_TOKEN="xoxb-YourBotTokenHere"

- export SLACK_APP_TOKEN="xapp-YourAppTokenHere"

### 3)Run the bot:

- go run main.go

##Usage
### 1)Invite the bot to your Slack workspace.
### 2)Use the command:

- /my yob is <year>

Example:

/my yob is 1990

# Project 2: Slack File Uploader

- This Go application uploads files to a specified Slack channel.
- This project showcases a Golang program that utilizes the Slack Go library to upload files to a specified Slack channel. The code allows you to set your Slack bot token and channel ID as environment variables. You can then upload files to the specified channel using the Slack API.

## Setup
### 1)Install dependencies:

- go get -u github.com/slack-go/slack

### 2)Set your Slack Bot Token and Channel ID as environment variables:

- export SLACK_BOT_TOKEN="xoxb-YourBotTokenHere"
- export CHANNEL_ID="YourChannelIDHere"

### 3)Run the application:

- go run main.go
## Usage
### 1)Set the fileArr variable with the paths of the files you want to upload.
### 2)Run the application, and it will upload the files to the specified Slack channel.
