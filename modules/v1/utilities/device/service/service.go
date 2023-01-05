package service

import (
	"ClearningPatternGO/modules/v1/utilities/device/models"
	"ClearningPatternGO/modules/v1/utilities/device/repository"
)

type Service interface {
	// ListProduct() ([]models.Product, error)
	ListDevice() ([]models.Device, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

// func (s *service) ListProduct() ([]models.Product, error) {
// 	allproduct, err := s.repository.ListProduct()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return allproduct, nil
// }

func (s *service) ListDevice() ([]models.Device, error) {
	allDevice, err := s.repository.ListDevice()
	if err != nil {
		return nil, err
	}
	return allDevice, nil
}
