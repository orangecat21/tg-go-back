# Telegram Mini App User Data Receiver

This is a Telegram bot that receives user data from Telegram Mini Apps.

## Features

- Receives user data from Telegram Mini Apps via HTTP endpoint
- Logs received user data to console
- Supports standard Telegram Mini App user data structure

## Endpoints

### POST `/user-data`
Accepts user data from Telegram Mini App in the following format:
```json
{
  "firstName": "string",
  "lastName": "string",
  "username": "string",
  "photoUrl": "string"
}
```

## How to Use

1. Start the bot
2. Send a POST request to `http://localhost:8080/user-data` with user data in JSON format
3. Check the console logs for received data

## Requirements

- Go 1.25+
- TELEGRAM_BOT_TOKEN environment variable set
