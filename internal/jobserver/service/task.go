package service

import (
	"job/internal/jobserver/pkg/model"
	"job/internal/jobserver/store"
)

type Task interface {
	Create(task *model.Task) error
	List(target string) (*[]model.Task, error)
	Update(task *model.Task) error
}

type TaskSrv struct {
	store store.Factory
}

func (t TaskSrv) Create(task *model.Task) error {
	return t.store.Task().Create(task)
}

func (t TaskSrv) List(target string) (*[]model.Task, error) {
	return t.store.Task().List(target)
}

func (t TaskSrv) Update(task *model.Task) error {
	return t.store.Task().Update(task)
}

func NewTaskSrv(store store.Factory) *TaskSrv {
	return &TaskSrv{
		store: store,
	}
}
