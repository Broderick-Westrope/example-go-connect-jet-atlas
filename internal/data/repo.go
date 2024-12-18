package data

import (
	"context"
	"database/sql"

	"github.com/Broderick-Westrope/eventurely-go/gen/dev/public/model"
)

type Respository interface {
	CreateUser(ctx context.Context, email, name string) (*model.Users, error)
	GetUserByID(ctx context.Context, id int64) (*model.Users, error)
	UpdateUser(ctx context.Context, id int64, email, name string) (*model.Users, error)
	DeleteUser(ctx context.Context, id int64) error
}

type repository_Impl struct {
	DB *sql.DB
}

func NewRepository(ctx context.Context, db *sql.DB) Respository {
	return &repository_Impl{
		DB: db,
	}
}
