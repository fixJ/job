package controller

import "job/internal/jobserver/service"

type TaskController struct {
	svc service.Factory
}

func NewTaskController(svc service.Factory) *TaskController {
	return &TaskController{
		svc: svc,
	}
}

// 添加任务
func (t *TaskController) Create() {

}

// 更新任务
func (t *TaskController) Update() {

}

// 拉取任务
func (t *TaskController) List() {

}
