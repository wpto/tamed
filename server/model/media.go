package model

type MediaMeta struct {
	Type   string `json:"type"` // ? 1 for vid 2 for img
	Mime   string `json:"mime"` // ? removed
	Audio  bool   `json:"audio"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// type MediaUrls struct {
// 	Sd     string  `json:"sd"`
// 	Hd     string  `json:"hd"`
// 	Poster string  `json:"poster"`
// 	Thumb  string  `json:"thumb"`
// 	Vthumb *string `json:"vthumb"`
// }

type MediaSocial struct {
	Likes int `json:"likes"`
	Views int `json:"views"`
}

func NewMediaSocial() *MediaSocial {
	return &MediaSocial{Likes: 0, Views: 0}
}

type Media struct {
	ID         string      `json:"id"`
	CreateTime string      `json:"create_time"`
	UserName   string      `json:"username"`
	Meta       MediaMeta   `json:"meta"`
	Social     MediaSocial `json:"social"`
	// Urls       MediaUrls   `json:"urls`
}

type Art struct {
	ID         string  `json:"id"`
	CreateTime string  `json:"create_time"`
	UserName   string  `json:"username"`
	Media      []Media `json:"media"`
}

type User struct {
	UserName string `json:"username"`
	Profile  struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		CreateTime  int    `json:"create_date"`
		PictureUrl  string `json:"picture_url"`
		Thumb       struct {
			Id      string `json:"id"`
			Poster  string `json:"poster"`
			Preview string `json:"preview"`
		} `json:"thumb"`
	} `json:"profile"`
	Social struct {
		Arts      int `json:"arts"`
		Followers int `json:"followers"`
		Following int `json:"following"`
		Views     int `json:"views"`
	} `json:"social"`
}
