package admin

import (
	"kcloudb1/internal/config"
	"time"
)

type Ingredient struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (p *Ingredient) TableName() string {
	return "one_ingredient"
}

func (p *Ingredient) Create() error {
	return config.DB.Create(p).Error
}

func (p *Ingredient) Get() error {
	return config.DB.Where("id = ?", p.ID).Find(p).Error
}

func (p *Ingredient) Update() error {
	return config.DB.Updates(p).Error
}

func (p *Ingredient) Delete() error {
	return config.DB.Where("id = ?", p.ID).Delete(p).Error
}

func (p *Ingredient) GetAll() ([]Ingredient, error) {
	var products []Ingredient
	return products, config.DB.Order("created_at desc").Find(&products).Error
}
