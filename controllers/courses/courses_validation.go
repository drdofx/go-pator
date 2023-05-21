package courses

type CreateCourseValidation struct {
	CourseName  string `json:"course_name" binding:"required"`
	CourseProdi string `json:"course_prodi" binding:"required"`
}

type UpdateCourseValidation struct {
	CourseName  string `json:"course_name"`
	CourseProdi string `json:"course_prodi"`
}
