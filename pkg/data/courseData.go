package data

type Course struct {
	ID         int64  `json:"id"`
	CourseName string `json:"name"`
	CourseCode string `json:"course_code"`
	Enrollment []struct {
		Type string `json:"type"`
	} `json:"enrollment"`
}
