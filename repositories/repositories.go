package repositories

import (
	"database/sql"
)

type IRepository interface {
}

type repository struct {
	db *sql.DB
}

func NewRepositorie(db *sql.DB) IRepository {
	return &repository{db: db}
}

// func (r *repository) {

// }
