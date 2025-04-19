package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env") // Loads .env into environment variables
    if err != nil {
			fmt.Printf("Error loading .env file: %s", err)
    }

	apiKey := os.Getenv("API_KEY")
  client := anthropic.NewClient(option.WithAPIKey(apiKey))

  scanner := bufio.NewScanner(os.Stdin)
  getUserMessage := func () (string, bool) {
    if !scanner.Scan(){
      return "", false
    }
    return scanner.Text(), true
  }

  agent := NewAgent(&client, getUserMessage)
  err = agent.Run(context.TODO())
  if err != nil {
    fmt.Printf("Error: %s\n", err.Error())
  }
}

func NewAgent(client *anthropic.Client, getUserMessage func() (string, bool)) *Agent {
  return &Agent{
    client: client,
    getUserMessage: getUserMessage,
  }
}

type Agent struct {
  client *anthropic.Client
  getUserMessage func() (string, bool)
}

func (a *Agent) Run(ctx context.Context) error {
	conversion := []anthropic.MessageParam{}

	fmt.Println("Chat with Claude (use 'ctrl-c' to quit)")

	for{
		fmt.Print("\u001b[94mYou\u001b[0m: ")
		userInput, ok := a.getUserMessage()
		if !ok {
			break
		}
		userMessage := anthropic.NewUserMessage(anthropic.NewTextBlock(userInput))
		conversion = append(conversion, userMessage)

		message, err := a.runInference(ctx, conversion)
		if err != nil {
			return err
		}
		conversion = append(conversion, message.ToParam())

		for _, content := range message.Content {
			switch content.Type {
			case "text":
				fmt.Printf("\u001b[92mClaude\u001b[0m: %s\n", content.Text)
			}
		}
	}

	return nil
}

func (a *Agent) runInference(ctx context.Context, conversion []anthropic.MessageParam) (*anthropic.Message, error) {
	message, err := a.client.Messages.New(ctx, anthropic.MessageNewParams{
		Model: anthropic.ModelClaude3_7SonnetLatest,
		MaxTokens: int64(1024),
		Messages: conversion,
	})
	return message, err
}