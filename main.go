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

	tools := []ToolDefinition{ReadFileDefinition, ListFileDefinition}
  agent := NewAgent(&client, getUserMessage, tools)
  err = agent.Run(context.TODO())
  if err != nil {
    fmt.Printf("Error: %s\n", err.Error())
  }
}