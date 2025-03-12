package cmd

import (
	"fmt"
	"log"
	"os"
)

func TextFile(text string) {
	err := os.WriteFile("studyPlan.txt", []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing to file")
		log.Fatal(err)
		return
	}

	fmt.Println("Successfully created file")
}
