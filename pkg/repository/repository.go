package repository

import (
	"github.com/jmoiron/sqlx"
)

type CreateUser interface {
}

type GetUserByID interface {
}

type DeleteUser interface {
}

type UpdateUser interface {
}

type UpdateAuth interface {
}

// Repository transition to postgres db
type Repository struct {
	CreateUser
	GetUserByID
	DeleteUser
	UpdateUser
	UpdateAuth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
