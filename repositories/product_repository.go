package repositories

import (
	"errors"
	"manajemen-product/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(id int, product models.Product) (models.Product, error)
	Delete(id int) error
	FindById(id int) (models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// Ambil semua data
func (r *productRepository) GetAll() ([]models.Product, error) {
	var product []models.Product

	err := r.db.Order("id desc").Find(&product).Error

	return product, err
}

func (r *productRepository) Create(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *productRepository) Update(id int, product models.Product) (models.Product, error) {
	var existingProduct models.Product

	if err := r.db.First(&existingProduct, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, errors.New("Data tidak ditemukan")
		}
		return product, err
	}

	err := r.db.Model(&existingProduct).Updates(product).Error

	if err != nil {
		return existingProduct, err
	}

	return existingProduct, nil
}

func (r *productRepository) Delete(id int) error {
	result := r.db.Delete(&models.Product{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Data tidak ditemukan")
	}

	return nil
}

func (r *productRepository) FindById(id int) (models.Product, error) {
	var product models.Product

	err := r.db.First(&product, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, errors.New("Data tidak ditemukan")
		}
		return product, err
	}

	return product, nil
}
