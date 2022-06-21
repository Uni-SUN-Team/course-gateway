package model

type CourseContent struct {
	Id           int64            `json:"id"`
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Content      string           `json:"content"`
	VideoContent Video            `json:"video_content"`
	VideoPreview Video            `json:"video_preview"`
	File         []ThumbnailLarge `json:"file"`
	Preview      bool             `json:"preview"`
	Component    string           `json:"__component"`
}
