package data

import (
	"context"
	"time"

	"github.com/Broderick-Westrope/eventurely-go/gen/dev/public/model"
	"github.com/Broderick-Westrope/eventurely-go/gen/dev/public/table"
	"github.com/go-jet/jet/v2/postgres"
)

func (repo *repository_Impl) CreateUser(ctx context.Context, email string, name string) (*model.Users, error) {
	user := &model.Users{
		ID:        0,
		Email:     email,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := table.Users.
		INSERT(table.Users.AllColumns.Except(table.Users.ID)).
		MODEL(user).
		RETURNING(table.Users.AllColumns).
		QueryContext(ctx, repo.DB, user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (repo *repository_Impl) GetUserByID(ctx context.Context, id int64) (*model.Users, error) {
	user := &model.Users{}

	err := table.Users.
		SELECT(table.Users.AllColumns).
		WHERE(table.Users.ID.EQ(postgres.Int(id))).
		QueryContext(ctx, repo.DB, user)
	if err != nil {
		if isNotFoundError(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return user, err
}

func (repo *repository_Impl) UpdateUser(ctx context.Context, id int64, email string, name string) (*model.Users, error) {
	user := &model.Users{
		ID:        id,
		Email:     email,
		Name:      name,
		CreatedAt: time.Time{},
		UpdatedAt: time.Now(),
	}

	err := table.Users.
		UPDATE(table.Users.AllColumns.
			Except(
				table.Users.ID,
				table.Users.CreatedAt,
			)).
		MODEL(user).
		WHERE(table.Users.ID.EQ(postgres.Int(id))).
		RETURNING(table.Users.AllColumns).
		QueryContext(ctx, repo.DB, user)
	if err != nil {
		if isNotFoundError(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return user, err
}

func (repo *repository_Impl) DeleteUser(ctx context.Context, id int64) error {
	_, err := table.Users.
		DELETE().
		WHERE(table.Users.ID.EQ(postgres.Int(id))).
		ExecContext(ctx, repo.DB)
	if err != nil {
		if isNotFoundError(err) {
			return ErrNotFound
		}
		return err
	}
	return nil
}
