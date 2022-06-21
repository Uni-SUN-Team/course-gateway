package model

type ThumbnailSmall struct {
	Id                int64                `json:"id"`
	Name              string               `json:"name"`
	AlternativeText   string               `json:"alternativeText"`
	Caption           string               `json:"caption"`
	Width             int64                `json:"width"`
	Height            int64                `json:"height"`
	Formats           ThumbnailSmallFormat `json:"formats"`
	Hash              string               `json:"hash"`
	Ext               string               `json:"ext"`
	Mime              string               `json:"mime"`
	Size              float64              `json:"size"`
	Url               string               `json:"url"`
	PreviewUrl        string               `json:"previewUrl"`
	Provider          string               `json:"provider"`
	Provider_metadata string               `json:"provider_metadata"`
	CreatedAt         string               `json:"createdAt"`
	UpdatedAt         string               `json:"updatedAt"`
}

type ThumbnailLarge struct {
	Id                int64                `json:"id"`
	Name              string               `json:"name"`
	AlternativeText   string               `json:"alternativeText"`
	Caption           string               `json:"caption"`
	Width             int64                `json:"width"`
	Height            int64                `json:"height"`
	Formats           ThumbnailLargeFormat `json:"formats"`
	Hash              string               `json:"hash"`
	Ext               string               `json:"ext"`
	Mime              string               `json:"mime"`
	Size              float64              `json:"size"`
	Url               string               `json:"url"`
	PreviewUrl        string               `json:"previewUrl"`
	Provider          string               `json:"provider"`
	Provider_metadata string               `json:"provider_metadata"`
	CreatedAt         string               `json:"createdAt"`
	UpdatedAt         string               `json:"updatedAt"`
}

type ThumbnailSmallFormat struct {
	Xsmall    FormatAttribute `json:"xsmall"`
	Thumbnail FormatAttribute `json:"thumbnail"`
}

type ThumbnailLargeFormat struct {
	Large     FormatAttribute `json:"large"`
	Small     FormatAttribute `json:"small"`
	Medium    FormatAttribute `json:"medium"`
	Xsmall    FormatAttribute `json:"xsmall"`
	Thumbnail FormatAttribute `json:"thumbnail"`
}

type FormatAttribute struct {
	Ext    string  `json:"ext"`
	Url    string  `json:"url"`
	Hash   string  `json:"hash"`
	Mime   string  `json:"mime"`
	Name   string  `json:"name"`
	Path   string  `json:"path"`
	Size   float64 `json:"size"`
	Width  int64   `json:"width"`
	Height int64   `json:"height"`
}
