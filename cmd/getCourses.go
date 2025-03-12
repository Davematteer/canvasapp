package cmd

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Davematteer/canvasapp/pkg/data"
)

const baseURL = "https://ashesi.instructure.com/api/v1"

func GetCourses() []data.Course {
	token := os.Getenv("CANVASTOKEN")
	client := &http.Client{}
	endpoints := []string{
		"/courses",
		"/users/self/profile",
		"/courses?enrollment_state=active",
		"courses?as_user_id=self",
		"users/:user_id/courses",
	}
	request, _ := http.NewRequest("GET", baseURL+endpoints[2], nil)

	// customize the request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-type", "application/json")

	response, err := client.Do(request)

	respBody := response.Body

	defer respBody.Close()

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(respBody)
	if err != nil {
		log.Fatal(err)
	}

	courseJson := string(body)

	// Parsing the json data
	var courses []data.Course
	err = json.Unmarshal([]byte(courseJson), &courses)
	if err != nil {
		log.Fatal(err)
	}

	return courses
}
