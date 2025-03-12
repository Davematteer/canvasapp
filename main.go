package main

import (
	"fmt"

	"github.com/Davematteer/canvasapp/cmd"
)

const baseURL = "https://ashesi.instructure.com/api/v1"

func main() {
	response := cmd.GroqChat().Choices
	r1 := ""
	if len(response) > 0 {
		result := response[0].Message.Content
		r1 += result
	} else {
		fmt.Println("No response found")
	}

	cmd.TextFile(r1)
	cmd.ConvertToPDF()
	cmd.SendEmail()
}
