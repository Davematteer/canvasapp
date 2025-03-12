package cmd

import (
	"encoding/json"
	"log"
)

func GetStructuredJson() string {
	jsonData, err := json.MarshalIndent(GetCourses(), "", " ")
	if err != nil {
		log.Fatal(err)
	}

	return string(jsonData)
}
