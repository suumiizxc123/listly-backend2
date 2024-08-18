package core

type City struct {
	UID    string `json:"uid"`
	NameMN string `json:"name_mn"`
	NameEN string `json:"name_en"`
}

func (c *City) TableName() string {
	return "cities"
}
