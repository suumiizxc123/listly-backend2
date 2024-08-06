package admin

import (
	"kcloudb1/internal/config"
	"time"

	"gorm.io/gorm/clause"
)

type Product struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	SalePrice   float32   `json:"sale_price"`
	CreatedAt   time.Time `json:"created_at"`
}

func (p *Product) TableName() string {
	return "one_product"
}

func (p *Product) Create() error {
	return config.DB.Create(p).Error
}

func (p *Product) Get() error {
	return config.DB.Where("id = ?", p.ID).Find(p).Error
}

func (p *Product) Update() error {
	return config.DB.Updates(p).Error
}

func (p *Product) Delete() error {
	return config.DB.Where("id = ?", p.ID).Delete(p).Error
}

func (p *Product) GetAll() ([]Product, error) {
	var products []Product
	return products, config.DB.Order("created_at desc").Find(&products).Error
}

type ProductIngredient struct {
	ID           int64     `json:"id"`
	ProductID    int64     `json:"product_id"`
	IngredientID int64     `json:"ingredient_id"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
}

func (p *ProductIngredient) TableName() string {
	return "one_product_ingredient"
}

func (p *ProductIngredient) Create() error {
	return config.DB.Create(p).Error
}

func (p *ProductIngredient) Get() error {
	return config.DB.Where("id = ?", p.ID).Find(p).Error
}

func (p *ProductIngredient) Update() error {
	return config.DB.Updates(p).Error
}

func (p *ProductIngredient) Delete() error {
	return config.DB.Where("id = ?", p.ID).Delete(p).Error
}

func (p *ProductIngredient) GetAll() ([]ProductIngredient, error) {
	var products []ProductIngredient
	return products, config.DB.Order("created_at desc").Find(&products).Error
}

type ProductImage struct {
	ID        int64     `json:"id"`
	ProductID int64     `json:"product_id"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *ProductImage) TableName() string {
	return "one_product_image"
}

func (p *ProductImage) Create() error {
	return config.DB.Create(p).Error
}

func (p *ProductImage) Get() error {
	return config.DB.Where("id = ?", p.ID).Find(p).Error
}

func (p *ProductImage) Update() error {
	return config.DB.Updates(p).Error
}

func (p *ProductImage) Delete() error {
	return config.DB.Where("id = ?", p.ID).Delete(p).Error
}

func (p *ProductImage) GetAll() ([]ProductImage, error) {
	var products []ProductImage
	return products, config.DB.Order("created_at desc").Find(&products).Error
}

type ProductExtend struct {
	ID          int64               `json:"id"`
	Title       string              `json:"title"`
	Subtitle    string              `json:"subtitle"`
	Description string              `json:"description"`
	Price       float32             `json:"price"`
	SalePrice   float32             `json:"sale_price"`
	Images      []ProductImage      `json:"images" gorm:"foreignKey:ProductID;references:ID"`
	Ingredients []ProductIngredient `json:"ingredients" gorm:"foreignKey:ProductID;references:ID"`
	CreatedAt   time.Time           `json:"created_at"`
}

type ProductExtend2 struct {
	ID          int64               `json:"id"`
	Title       string              `json:"title"`
	Subtitle    string              `json:"subtitle"`
	Description string              `json:"description"`
	Price       float32             `json:"price"`
	SalePrice   float32             `json:"sale_price"`
	Images      []ProductImage      `json:"images" gorm:"foreignKey:ProductID;references:ID"`
	Ingredients []ProductIngredient `json:"ingredients" gorm:"foreignKey:ProductID;references:ID"`
	CreatedAt   time.Time           `json:"created_at"`
	Images2     []string            `json:"images2"`
}

func (p *ProductExtend) TableName() string {
	return "one_product"
}

func (p *ProductExtend) Get() error {
	return config.DB.Where("id = ?", p.ID).Preload(clause.Associations).Find(p).Error
}

func (p *ProductExtend) GetAll() ([]ProductExtend, error) {
	var products []ProductExtend
	return products, config.DB.Order("created_at desc").Preload(clause.Associations).Find(&products).Error
}

type ProductImageInput struct {
	Image string `json:"image"`
}

type ProductIngredientInput struct {
	IngredientID int64  `json:"ingredient_id"`
	Description  string `json:"description"`
}
type ProductInput struct {
	Title       string                   `json:"title"`
	Subtitle    string                   `json:"subtitle"`
	Description string                   `json:"description"`
	Price       float32                  `json:"price"`
	SalePrice   float32                  `json:"sale_price"`
	Images      []ProductImageInput      `json:"images"`
	Ingredients []ProductIngredientInput `json:"ingredients"`
}
