package usecase

import (
	"context"
	"github.com/sviridoves/GB_Backend_Level2/lesson1/internal/pkg/models"
	"github.com/sviridoves/GB_Backend_Level2/lesson1/internal/pkg/user"
)

type usecase struct {
	repo user.Repository
}

func New(repo user.Repository) user.Usecase {
	result := usecase{
		repo: repo,
	}
	return result
}

func (u usecase) Create(ctx context.Context, user *models.User) error {
	panic("implement me")
}

func (u usecase) Add(ctx context.Context, user *models.User) error {
	panic("implement me")
}

func (u usecase) Delete(ctx context.Context, user *models.User) error {
	panic("implement me")
}

func (u usecase) Find(ctx context.Context, user *models.User) error {
	panic("implement me")
}
