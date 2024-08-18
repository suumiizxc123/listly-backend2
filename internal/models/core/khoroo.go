package core

type Khoroo struct {
	UID         string `json:"uid"`
	DistrictUID string `json:"district_uid"`
	Name        string `json:"name"`
}

func (k *Khoroo) TableName() string {
	return "khoroo"
}
