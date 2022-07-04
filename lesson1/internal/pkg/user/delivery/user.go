package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/sviridoves/GB_Backend_Level2/lesson1/internal/pkg/user"
)

type delivery struct {
	users user.Usecase
}

func New(users user.Usecase) user.Delivery {
	result := delivery{
		users: users,
	}
	return result
}

func (d delivery) Create() gin.HandlerFunc {
	panic("implement me")
}

func (d delivery) Add() gin.HandlerFunc {
	panic("implement me")
}

func (d delivery) Delete() gin.HandlerFunc {
	panic("implement me")
}

func (d delivery) Find() gin.HandlerFunc {
	panic("implement me")
}
