package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

// we use the github.com/ollama/ollama/api pakage to create a client and send requests to the Ollama API.
// *** github.com/ollama/ollama/api ***
func main() {
	// 1. create a ollama client
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Prepared the request for the model.
	req := &api.GenerateRequest{
		Model:  "llama3.2",
		Prompt: "what his name?",
		Stream: nil,
	}

	ctx := context.Background()

	err = client.Generate(ctx, req, func(resp api.GenerateResponse) error {
		fmt.Print(resp.Response)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
}
