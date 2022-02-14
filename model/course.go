package model

type Course struct {
	CourseID   int    `xorm:"course_id INTEGER pk"`
	CourseName string `xorm:"TEXT"`
}

func init() {

}

func AddCourse(id int, name string) {
	c := new(Course)

	c.CourseID = id
	c.CourseName = name
	_, err := engine.InsertOne(c)
	if err != nil {
		return
	}
}

func CourseQuery(id int) *Course {
	c := new(Course)
	_, err := engine.Where("course_id=?", id).Get(c)
	if err != nil {
		return nil
	}
	c.CourseID = id
	return c
}
