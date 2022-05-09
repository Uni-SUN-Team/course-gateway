package model

type Course struct {
	Data CoursesData       `json:"data"`
	Meta CoursesPagination `json:"meta"`
}

type Courses struct {
	Data []CoursesData     `json:"data"`
	Meta CoursesPagination `json:"meta"`
}

type CoursesData struct {
	Id         int64             `json:"id"`
	Attributes CoursesAttributes `json:"attributes"`
}

type CoursesAttributes struct {
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Content       string     `json:"content"`
	Short_content string     `json:"short_content"`
	Price         float64    `json:"price"`
	CreatedAt     string     `json:"createdAt"`
	UpdatedAt     string     `json:"updatedAt"`
	PublishedAt   string     `json:"publishedAt"`
	Locale        string     `json:"locale"`
	Slug          string     `json:"slug"`
	Thumnail      Thumnail   `json:"thumnail"`
	Categories    Categories `json:"categories"`
	Advisors      Advisors   `json:"advisors"`
}

type CoursesPagination struct {
	Pagination CoursesPaginationContent `json:"pagination"`
}

type CoursesPaginationContent struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"pageSize"`
	PageCount int64 `json:"pageCount"`
	Total     int64 `json:"total"`
}
