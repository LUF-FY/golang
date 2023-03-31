package sqlite

import "database/sql"

type UserRepo struct {
	dbClient *sql.DB
}

func NewRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		dbClient: db,
	}
}
