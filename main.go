package main

import (
	"context"
	"fmt"
	"log"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func main() {

	// Point the OpenAI SDK to the local Ollama server.
	client := openai.NewClient(
		option.WithBaseURL("http://localhost:11434/v1"),
		option.WithAPIKey("ollama"),
	)

	ctx := context.Background()

	//message to send to the model
	resp, err := client.Chat.Completions.New(
		ctx,
		openai.ChatCompletionNewParams{
			Model: "llama3.2:latest",
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage("Tell me a joke about Go programming."),
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
