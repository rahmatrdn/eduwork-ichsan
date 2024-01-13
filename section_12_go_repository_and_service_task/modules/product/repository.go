package product

import (
	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		Create(product *Product) (*Product, error)
		GetAll() ([]Product, error)
		GetByID(id int) (*Product, error)
		Update(product *Product) (*Product, error)
		Delete(product *Product, id int) (*Product, error)
	}

	productRepository struct {
		DB *gorm.DB
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		DB: db,
	}
}

func (pr productRepository) Create(product *Product) (*Product, error) {
	if err := pr.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (pr productRepository) GetAll() ([]Product, error) {
	var products []Product
	if err := pr.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pr productRepository) GetByID(id int) (*Product, error) {
	var product Product
	if err := pr.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (pr productRepository) Update(product *Product) (*Product, error) {
	if err := pr.DB.Save(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (pr productRepository) Delete(product *Product, id int) (*Product, error) {
	if err := pr.DB.Delete(&product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}
