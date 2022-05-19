package model

type ResponseCourseAll struct {
	Error      error                    `json:"error"`
	Result     []CoursesDatas           `json:"result"`
	Pagination CoursesPaginationContent `json:"pagination"`
}

type ResponseCourse struct {
	Error  error       `json:"error"`
	Result CoursesData `json:"result"`
}
