package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"monolith/models"
	"monolith/query"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (s *productRepo) GetProductCtrl(page, limit int) ([]models.Category, error) {
    offset :=(page-1)*limit
	
	
	row, err := s.db.Query(query.CATEGORY_SQL, offset, limit)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error in GetProductCtrl: %w", err)
	}
	var categories []models.Category

	for row.Next() {
		var category models.Category
		row.Scan(
			&category.Id,
			&category.Name)
		categories = append(categories, category)
	}

	for i, category := range categories {
		rows, err := s.db.Query(query.PRODUCT_SQL, category.Id)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("error in GetProductCtrl: %w", err)
		}
		for rows.Next() {
			var product models.Product
			rows.Scan(
				&product.Id,
				&product.Name,
				&product.Date)
			categories[i].Products = append(categories[i].Products, product)

		}
	}
	return categories, nil

}
