package mysql

import (
	"gorm.io/gorm"
	"job/internal/jobserver/pkg/model"
)

type TaskStore struct {
	db *gorm.DB
}

func NewTaskStore(db *gorm.DB) *TaskStore {
	return &TaskStore{db: db}
}

func (t TaskStore) Get(target, tid string) (*model.Task, error) {
	r := &model.Task{}
	err := t.db.Where("target = ? and task_id = ?", target, tid).First(&r).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (t *TaskStore) Create(task *model.Task) error {
	return t.db.Create(&task).Error
}

func (t *TaskStore) List(target string) (*[]model.Task, error) {
	ret := &[]model.Task{}
	err := t.db.Where("target = ?", target).Find(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (t *TaskStore) Update(task *model.Task) error {
	return t.db.Save(&task).Error
}
