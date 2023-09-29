package platform

import (
	"database/sql"
	"fmt"
	"monolith/config"

	_ "github.com/lib/pq"
)

func DBConn() *sql.DB {
	cfg := config.Load()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}
