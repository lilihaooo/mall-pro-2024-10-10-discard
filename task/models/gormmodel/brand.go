package gormmodel

type Brand struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (*Brand) TableName() string {
	return "brands"
}
