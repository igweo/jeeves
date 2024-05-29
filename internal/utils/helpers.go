package utils

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

var openaiToken = os.Getenv("OPENAI_API_KEY")

func ChatCompletion(message string) string {
	if len(openaiToken) == 0 {
		return "'OPENAI_API_KEY' not set in ENV"
	}

	client := openai.NewClient(openaiToken)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Do not give ANY premable or any answer that includes triple back tick Markdown segments as it will be directly dumped into my terminal",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "Completion Error - Try again"
	}
	return resp.Choices[0].Message.Content
}
