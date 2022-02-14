package config

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type CheckQrCode struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
}

type GetCourseList struct {
	// 用户id
	ID int `json:"id"`
	// 课程id
	CourseID int `json:"course_id"`
}

type DoActive struct {
	// 用户id
	UserID int `json:"user_id"`
	// 课程id
	CourseID int `json:"course_id"`
}
