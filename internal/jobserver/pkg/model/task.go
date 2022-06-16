package model

import validator "gopkg.in/go-playground/validator.v9"

type Task struct {
	Base
	Name    string `json:"name" binding:"required" validate:"min=1 max=64" gorm:"column:name"`
	Target  string `json:"target" binding:"required" validate:"min=7 max15" gorm:"column:target"`
	Command string `json:"command" binding:"required" gorm:"column:command"`
	Status  int    `json:"status" binding:"required" gorm:"column:status"`   // 0 已创建未开始, 1 成功, -1 失败
	Cron    bool   `json:"cron" binding:"required" gorm:"column:cron"`       // true定时 false立即执行
	Runtime uint64 `json:"runtime" binding:"required" gorm:"column:runtime"` // 定时为时间戳, 立即执行为0
}

func (t *Task) TableName() string {
	return "task"
}

func (t *Task) Validate() error {
	return validator.New().Struct(t)
}
