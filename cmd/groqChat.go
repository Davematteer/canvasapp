package cmd

import (
	"log"
	"os"

	"github.com/jpoz/groq"
)

func GroqChat() *groq.ChatCompletion {
	groqKey := os.Getenv("GROQKEY")
	client := groq.NewClient(groq.WithAPIKey(groqKey))

	// jsonData := GetStructuredJson()

	response, err := client.CreateChatCompletion(groq.CompletionCreateParams{
		Model: "llama3-8b-8192",
		Messages: []groq.Message{
			{
				Role:    "user",
				Content: LoopCourses(GetCourses()),
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	return response
}
