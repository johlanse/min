package config

type Config struct {
	Study []struct {
		ID       int `json:"id"`
		CourseID int `json:"course_id"`
	} `json:"study"`
}
