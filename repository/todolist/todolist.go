package todolist

import (
	"github.com/natashameiliana/learnapi/dto/entity"
	"time"
)

type UseCase interface {
	WriteTask(data entity.Todolist) error
	ReadTask(data entity.Todolist) ([]entity.Todolist, error)
	UpdateTask(data entity.Todolist) error
	EraseTask(UserId string) error
}

type repository struct {
}

func (r *repository) WriteTask(data entity.Todolist) error {
	return nil
}

func (r *repository) ReadTask(data entity.Todolist) ([]entity.Todolist, error) {
	return []entity.Todolist{
		{
			ID:        "11111",
			Title:     "AYAY",
			Desc:      "YOI MAMEN",
			Done:      false,
			Category:  "PERSONAL",
			Priority:  1,
			DueDate:   "NTAR",
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	}, nil
}

func (r *repository) UpdateTask(data entity.Todolist) error {
	return nil
}

func (r *repository) EraseTask(UserId string) error {
	return nil
}

func NewRepository() UseCase {
	return &repository{}
}
