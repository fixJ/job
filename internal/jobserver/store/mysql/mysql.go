package mysql

import (
	"fmt"
	"gorm.io/gorm"
	"job/internal/jobserver/store"
	"job/pkg/db"
	"sync"
	"time"
)

var (
	d    *gorm.DB
	once sync.Once
)

type Store struct {
	db *gorm.DB
}

func (s Store) Task() store.Task {
	return NewTaskStore(s.db)
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func GetMySQLInsOr(host, username, password, database string) (*gorm.DB, error) {
	var err error
	c := db.Config{
		Host:                  host,
		Username:              username,
		Password:              password,
		Database:              database,
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: 10 * time.Second,
	}
	once.Do(func() {
		d, err = db.New(&c)
	})
	if d == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", d, err)
	}
	return d, nil
}
