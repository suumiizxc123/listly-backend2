package core

type District struct {
	UID     string `json:"uid"`
	CityUID string `json:"city_uid"`
	NameMn  string `json:"name_mn"`
	NameEn  string `json:"name_en"`
}

func (c *District) TableName() string {
	return "districts"
}
