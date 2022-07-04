package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sviridoves/GB_Backend_Level2/lesson1/internal/pkg/models"
)

type Delivery interface {
	Create() gin.HandlerFunc
	Add() gin.HandlerFunc
	Delete() gin.HandlerFunc
	Find() gin.HandlerFunc
}

type Usecase interface {
	Create(ctx context.Context, user *models.User) error
	Add(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, user *models.User) error
	Find(ctx context.Context, user *models.User) error
}

type Repository interface {
	Create(ctx context.Context, user *models.User) error
	Add(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, user *models.User) error
	Find(ctx context.Context, user *models.User) error
}
