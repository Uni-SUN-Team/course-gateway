package model

type Advisors struct {
	Data []AdvisorsData `json:"data"`
}

type AdvisorsData struct {
	Id         int64              `json:"id"`
	Attributes AdvisorsAttributes `json:"attributes"`
}

type AdvisorsAttributes struct {
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	Telephone     string `json:"telephone"`
	Dob           string `json:"dob"`
	Content       string `json:"content"`
	Short_content string `json:"short_content"`
	Active        bool   `json:"active"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	PublishedAt   string `json:"publishedAt"`
	Locale        string `json:"locale"`
	Slug          string `json:"slug"`
}
