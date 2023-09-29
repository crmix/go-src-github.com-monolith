package storage

import (
	"database/sql"
	"monolith/models"
	"monolith/storage/postgres"
)

type ProductI interface {
	GetProductCtrl(page, limit int) ([]models.Category, error)
}

func NewProductRepo(db *sql.DB) ProductI {
	return postgres.NewProductRepo(db)
}
