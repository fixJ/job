package mysql

import (
	"gorm.io/gorm"
	"job/internal/jobserver/store"
)

type Store struct {
	db *gorm.DB
}

func (s Store) task() store.Task {
	return NewTaskStore(s.db)
}

func NewStore() {

}
