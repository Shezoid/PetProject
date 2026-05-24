package pkg

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository struct {
	Db *sql.DB
}

func NewRepository(connectionString string) *Repository {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	return &Repository{Db: db}
}
