# Go AI Agent

A simple command-line chat interface for Claude AI using the Anthropic Go SDK.

## Prerequisites

- Go 1.24.2 or later
- An Anthropic API key ([get one here](https://console.anthropic.com/))

## Setup

1. Clone the repository:

```bash
git clone <your-repo-url>
cd <your-repo-directory>
```

2. Create a `.env` file in the project root:

```bash
touch .env
```

3. Add your Anthropic API key to the `.env` file:

```
API_KEY=sk-ant-xxxx...  # Replace with your actual API key
```

4. Install dependencies:

```bash
go mod download
```

## Running the Application

Run the chat interface:

```bash
go run main.go
```

The application will start a chat session where you can interact with Claude. Type your messages and press Enter to send. Use Ctrl+C to exit the chat.

## Features

- Interactive command-line chat interface
- Colored output for better readability
- Conversation history maintained during the session
- Environment-based configuration
- Graceful error handling

## Security Note

Never commit your `.env` file or share your API key. The `.env` file is included in `.gitignore` by default.
