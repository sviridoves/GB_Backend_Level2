package sqlite

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sviridoves/GB_Backend_Level2/lesson1/internal/pkg/environment"
	"github.com/sviridoves/GB_Backend_Level2/lesson1/internal/pkg/models"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) environment.Repository {
	resp := repository{db: db}

	return resp
}

func (r repository) Create(ctx context.Context, user *models.Environment) error {
	panic("implement me")
}

func (r repository) Add(ctx context.Context, user *models.Environment) error {
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, user *models.Environment) error {
	panic("implement me")
}

func (r repository) Find(ctx context.Context, user *models.Environment) error {
	panic("implement me")
}
