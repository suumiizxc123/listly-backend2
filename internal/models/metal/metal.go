package metal

import "kcloudb1/internal/config"

type Metal struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Metric string `json:"metric"`
}

func (m *Metal) TableName() string {
	return "one_metal"
}

func (m *Metal) Create() error {
	return config.DB.Create(m).Error
}

func (m *Metal) Update() error {
	return config.DB.Updates(m).Error
}

func (m *Metal) Delete(id any) error {
	return config.DB.Delete(m, id).Error
}

func (m *Metal) Get(id any) error {
	return config.DB.First(m, id).Error
}

func (m *Metal) GetAll() ([]Metal, error) {
	var metals []Metal
	err := config.DB.Find(&metals).Error
	return metals, err
}
