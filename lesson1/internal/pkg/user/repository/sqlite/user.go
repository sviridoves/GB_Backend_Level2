package sqlite

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sviridoves/GB_Backend_Level2/lesson1/internal/pkg/models"
	"github.com/sviridoves/GB_Backend_Level2/lesson1/internal/pkg/user"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) user.Repository {
	rep := repository{
		db: db,
	}
	return rep
}

func (r repository) Create(ctx context.Context, user *models.User) error {
	panic("implement me")
}

func (r repository) Add(ctx context.Context, user *models.User) error {
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, user *models.User) error {
	panic("implement me")
}

func (r repository) Find(ctx context.Context, user *models.User) error {
	panic("implement me")
}
