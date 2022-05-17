package model

type CourseHighLight struct {
	Id                int64   `json:"id"`
	Name              string  `json:"name"`
	AlternativeText   string  `json:"alternativeText"`
	Caption           string  `json:"caption"`
	Width             int64   `json:"width"`
	Height            int64   `json:"height"`
	Hash              string  `json:"hash"`
	Ext               string  `json:"ext"`
	Mime              string  `json:"mime"`
	Size              float64 `json:"size"`
	Url               string  `json:"url"`
	PreviewUrl        string  `json:"previewUrl"`
	Provider          string  `json:"provider"`
	Provider_metadata string  `json:"provider_metadata"`
	CreatedAt         string  `json:"createdAt"`
	UpdatedAt         string  `json:"updatedAt"`
}
