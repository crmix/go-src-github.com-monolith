package models

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Products []Product
}
type Product struct {
	Id    int       `json:"id"`
	Name  string    `json:"name"`
	Date  string    `json:"date"`
}
