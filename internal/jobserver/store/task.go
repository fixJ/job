package store

import "job/internal/jobserver/pkg/model"

type Task interface {
	Get(target, tid string) (*model.Task, error)
	Create(task *model.Task) error
	List(target string) (*[]model.Task, error)
	Update(task *model.Task) error
}
