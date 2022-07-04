package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/sviridoves/GB_Backend_Level2/lesson1/internal/pkg/environment"
)

type delivery struct {
	users environment.Usecase
}

func New(env environment.Usecase) environment.Delivery {
	resp := delivery{users: env}
	return resp
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
