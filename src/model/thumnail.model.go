package model

type Thumnail struct {
	Id                int64                        `json:"id"`
	Name              string                       `json:"name"`
	AlternativeText   string                       `json:"alternativeText"`
	Caption           string                       `json:"caption"`
	Width             int64                        `json:"width"`
	Height            int64                        `json:"height"`
	Hash              string                       `json:"hash"`
	Ext               string                       `json:"ext"`
	Mime              string                       `json:"mime"`
	Size              float64                      `json:"size"`
	Url               string                       `json:"url"`
	PreviewUrl        string                       `json:"previewUrl"`
	Provider          string                       `json:"provider"`
	Provider_metadata string                       `json:"provider_metadata"`
	CreatedAt         string                       `json:"createdAt"`
	UpdatedAt         string                       `json:"updatedAt"`
	Formats           ThumnailDataAttributesFormat `json:"formats"`
}

type ThumnailDataAttributesFormat struct {
	Large    ThumnailDataAttributesFormatLarge     `json:"large"`
	Small    ThumnailDataAttributesFormatSmall     `json:"small"`
	Medium   ThumnailDataAttributesFormatMedium    `json:"medium"`
	Xsmall   ThumnailDataAttributesFormatXsmall    `json:"xsmall"`
	Thumnail ThumnailDataAttributesFormatThumbnail `json:"thumbnail"`
}

type ThumnailAttribute struct {
	Id                int64                                 `json:"id"`
	Name              string                                `json:"name"`
	AlternativeText   string                                `json:"alternativeText"`
	Caption           string                                `json:"caption"`
	Width             int64                                 `json:"width"`
	Height            int64                                 `json:"height"`
	Hash              string                                `json:"hash"`
	Ext               string                                `json:"ext"`
	Mime              string                                `json:"mime"`
	Size              float64                               `json:"size"`
	Url               string                                `json:"url"`
	PreviewUrl        string                                `json:"previewUrl"`
	Provider          string                                `json:"provider"`
	Provider_metadata string                                `json:"provider_metadata"`
	CreatedAt         string                                `json:"createdAt"`
	UpdatedAt         string                                `json:"updatedAt"`
	Formats           ThumnailDataAttributesFormatAttribute `json:"formats"`
}

type ThumnailDataAttributesFormatAttribute struct {
	Xsmall   ThumnailDataAttributesFormatXsmall    `json:"xsmall"`
	Thumnail ThumnailDataAttributesFormatThumbnail `json:"thumbnail"`
}

type ThumnailDataAttributesFormatLarge struct {
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

type ThumnailDataAttributesFormatSmall struct {
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

type ThumnailDataAttributesFormatMedium struct {
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

type ThumnailDataAttributesFormatXsmall struct {
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

type ThumnailDataAttributesFormatThumbnail struct {
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
