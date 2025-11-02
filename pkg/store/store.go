package store

import "database/sql"

type UserStore interface {
	GetUser(id int) (string, error)
}

type UserStoreImpl struct {
	db *sql.DB
}

func (store *UserStoreImpl) GetUser(id int) (string, error) {
	return "", nil
}
