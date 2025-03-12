package cmd

import (
	"fmt"

	"github.com/Davematteer/canvasapp/pkg/data"
)

func formatprompt(course data.Course) string {
	prompt := fmt.Sprintf(
		"You are an AI tutor. Given the course structure below, generate a study plan that includes:\n"+
			"- A short summary of each module.\n"+
			"- A list of key topics to focus on.\n"+
			"- Practice questions.\n\n"+
			"Course: %s\nCourseId: %d\nCourseCode: %s\n",
		course.CourseName,
		course.ID,
		course.CourseCode,
	)

	return prompt
}

func LoopCourses(cl []data.Course) string {
	loopedString := ""

	for _, c := range cl {
		loopedString += formatprompt(c)
	}

	return loopedString
}
