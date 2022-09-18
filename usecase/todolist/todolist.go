package todolist

import "github.com/natashameiliana/learnapi/dto/model"

type UseCase interface {
	SaveTask(req model.CreateTaskRequest) (model.CreateTaskResponse, error)
	MyTask(req model.GetTaskRequest) (model.GetTaskResponse, error)
	EditTask(req model.UpdateTaskRequest) (model.CreateTaskResponse, error)
	RemoveTask(req model.DeleteTaskRequest) (model.DeleteTaskResponse, error)
}

type usecase struct {
}

func (u *usecase) SaveTask(req model.CreateTaskRequest) (model.CreateTaskResponse, error) {
	panic("implement me")
}

func (u *usecase) MyTask(req model.GetTaskRequest) (model.GetTaskResponse, error) {
	panic("implement me")
}

func (u *usecase) EditTask(req model.UpdateTaskRequest) (model.CreateTaskResponse, error) {
	panic("implement me")
}

func (u *usecase) RemoveTask(req model.DeleteTaskRequest) (model.DeleteTaskResponse, error) {
	panic("implement me")
}

func NewUserCase() UseCase {
	return &usecase{}
}
