package gormmodel

type Image struct {
	ID  uint   `json:"id"`
	Url string `json:"url"`
}

func (Image) TableName() string {
	return "images"
}
