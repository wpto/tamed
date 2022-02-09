package model

type MediaMeta struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Mime   string `json:"mime"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
