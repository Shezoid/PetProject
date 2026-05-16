package pkg

import "database/sql"

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
