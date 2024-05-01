package common

type Language struct {
	ID       int64  `json:"ID" gorm:"primary_key"`
	Language string `json:"language"`
}

func (c *Language) TableName() string {
	return "language"
}
