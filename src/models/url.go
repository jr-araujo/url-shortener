package models

type InputUrl struct {
	Url string
}

type ShortenUrlOutput struct {
	Code       string
	ShortenUrl string
}

type ShortenUrl struct {
	Id            int32  `json:"id" gorm:"primary"`
	Code          string `json:"code" gorm:"code"`
	Original      string `json:"original" gorm:"original"`
	Access_number int32  `json:"access_number" gorm:"access_number"`
}

func (ShortenUrl) TableName() string {
	return "shorten_url"
}
