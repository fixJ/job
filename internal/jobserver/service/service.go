package service

import "job/internal/jobserver/store"

type Factory interface {
	Task() Task
}

type service struct {
	store store.Factory
}

func (s service) Task() Task {
	return NewTaskSrv(s.store)
}

func NewService(store store.Factory) *service {
	return &service{
		store: store,
	}
}
