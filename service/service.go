package service

import (
	"fmt"
	"log"
	"monolith/models"
	"monolith/storage"
)

type productServiceImpl struct {
	storage storage.ProductI
	logger log.Logger
}

func NewProductService(productI storage.ProductI, logger log.Logger) ProductService {
	return &productServiceImpl{
		storage: productI,
		logger: logger,
	}
}

type ProductService interface {
	GetProductCtrl(page, limit int) ([]models.Category, error)
}

func (c *productServiceImpl) GetProductCtrl(page, limit int) ([]models.Category, error) {
	res, err := c.storage.GetProductCtrl(page, limit)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error in getproductctrl %w", err)
	}

	return res, nil
}
