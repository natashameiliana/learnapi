package usecase

import "github.com/natashameiliana/learnapi/usecase/todolist"

type Pool struct {
	TodolistUsecase todolist.UseCase
}

func NewPoolUsecases(todolistUsecase todolist.UseCase) Pool {
	return Pool{
		TodolistUsecase: todolistUsecase,
	}
}
