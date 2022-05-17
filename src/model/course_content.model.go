package model

type CourseContent struct {
	Id           int64        `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Content      string       `json:"content"`
	VideoContent VideoContent `json:"video_content"`
	VideoPreview VideoPreview `json:"video_preview"`
	File         []Thumnail     `json:"file"`
}

type VideoContent struct {
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

type VideoPreview struct {
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
