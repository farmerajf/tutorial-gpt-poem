package main

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func main() {

	// Read user input
	var input string
	println("Enter a word to create a poem: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
	}

	// Create an instance of the OpenAI client
	c := openai.NewClient("YOUR_API_KEY")
	ctx := context.Background()

	// Create a chat completion request
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 100, // This is the number of tokens used to generate the poem
		// The higher the number, the longer the poem
		// This is a list initial messages to give the model context
		// The more messages, the more context the model has to generate the poem
		// In this case, we only need one message
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Write a short poem about a " + input + ":\n",
			},
		},
	}

	// Send the request to the API
	resp, err := c.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	// Print the poem
	fmt.Println(resp.Choices[0].Message.Content)
}
