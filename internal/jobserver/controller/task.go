package controller

import (
	"encoding/json"
	"fmt"
	"job/internal/jobserver/pkg/model"
	"job/internal/jobserver/service"
	"job/pkg/constant"
	"net/http"
)

type TaskController struct {
	svc service.Factory
}

func NewTaskController(svc service.Factory) *TaskController {
	return &TaskController{
		svc: svc,
	}
}

// 添加任务
// name 任务名
// target 目标主机ip
// command 执行命令
// cron 是否定时
// runtime 定时时间戳
func (t *TaskController) Create(w http.ResponseWriter, req *http.Request) {
	body := make([]byte, req.ContentLength)
	var cr CreateReq
	req.Body.Read(body)
	err := json.Unmarshal(body, &cr)
	if err != nil {
		resp, _ := json.Marshal(CommonResp{
			Code:    -1,
			Message: "request error",
		})
		w.Write(resp)
		return
	}
	fmt.Println(cr)
	task := model.Task{
		Name:    cr.Name,
		Target:  cr.Target,
		Command: cr.Command,
		Cron:    cr.Cron,
		Runtime: cr.Runtime,
		Status:  constant.TASKSTATUSSUCCESS,
	}
	err = t.svc.Task().Create(&task)
	if err != nil {
		resp, _ := json.Marshal(CommonResp{
			Code:    -2,
			Message: "create task error",
		})
		w.Write(resp)
		return
	}
	resp, _ := json.Marshal(CommonResp{
		Code:    0,
		Message: "ok",
	})
	w.Write(resp)
}

// 更新任务
func (t *TaskController) Update() {

}

// 拉取任务
func (t *TaskController) List() {

}
