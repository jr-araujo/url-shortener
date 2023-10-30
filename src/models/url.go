package models

type InputUrl struct {
	Url string
}

type ShortenUrl struct {
	Id            int32  `json:"id" gorm:"primary"`
	Code          string `json:"code"`
	Original_url  string `json:"original"`
	Shorten_url   string `json:"shorten_url"`
	Access_number int32  `json:"access_number"`
}

func (ShortenUrl) TableName() string {
	return "shorten_urls"
}
