package repository

import (
	"ClearningPatternGO/modules/v1/utilities/device/models"

	"gorm.io/gorm"
)

type Repository interface {
	// ListProduct() ([]models.Product, error)
	ListDevice() ([]models.Device, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) ListProduct() ([]models.Product, error) {
// 	var products []models.Product
// 	err := r.db.Find(&products).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return products, nil
// }

func (r *repository) ListDevice() ([]models.Device, error) {
	var device []models.Device
	err := r.db.Raw("SELECT * from device").Scan(&device).Error
	if err != nil {
		return nil, err
	}
	return device, nil
}
