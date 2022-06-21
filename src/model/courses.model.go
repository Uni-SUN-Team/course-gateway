package model

type Course struct {
	Data CoursesData       `json:"data"`
	Meta CoursesPagination `json:"meta"`
}

type CourseSlug struct {
	Data []CoursesData     `json:"data"`
	Meta CoursesPagination `json:"meta"`
}

type Courses struct {
	Data []CoursesDatas    `json:"data"`
	Meta CoursesPagination `json:"meta"`
}

type CoursesData struct {
	Id              int64           `json:"id"`
	Title           string          `json:"title"`
	Description     string          `json:"description"`
	Content         string          `json:"content"`
	Short_content   string          `json:"short_content"`
	CreatedAt       string          `json:"createdAt"`
	UpdatedAt       string          `json:"updatedAt"`
	PublishedAt     string          `json:"publishedAt"`
	Locale          string          `json:"locale"`
	Slug            string          `json:"slug"`
	Thumbnail       ThumbnailLarge  `json:"thumbnail"`
	Categories      []Categories    `json:"categories"`
	Advisors        []Advisors      `json:"advisors"`
	CourseHighLight Video           `json:"course_high_light"`
	Course          []CourseContent `json:"Course"`
	Rate            Rate            `json:"rate"`
}

type CoursesDatas struct {
	Id            int64          `json:"id"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	Content       string         `json:"content"`
	Short_content string         `json:"short_content"`
	CreatedAt     string         `json:"createdAt"`
	UpdatedAt     string         `json:"updatedAt"`
	PublishedAt   string         `json:"publishedAt"`
	Locale        string         `json:"locale"`
	Slug          string         `json:"slug"`
	Thumbnail     ThumbnailLarge `json:"thumbnail"`
	Advisors      []Advisors     `json:"advisors"`
	Rate          Rate           `json:"rate"`
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
