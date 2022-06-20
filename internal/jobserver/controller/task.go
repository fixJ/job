package controller

import (
	"encoding/json"
	"fmt"
	"job/internal/jobserver/manager"
	"job/internal/jobserver/pkg/model"
	"job/internal/jobserver/pkg/types"
	"job/internal/jobserver/service"
	"job/pkg/constant"
	"net/http"
	"time"
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
	var cr types.CreateReq
	req.Body.Read(body)
	err := json.Unmarshal(body, &cr)
	if err != nil {
		resp, _ := json.Marshal(types.CommonResp{
			Code:    -1,
			Message: "request error",
		})
		w.Write(resp)
		return
	}
	m, err := manager.GetManagerOr()
	if !m.IsLive(cr.Target) {
		resp, _ := json.Marshal(types.CommonResp{
			Code:    -1,
			Message: "can't create task, because the target ip not live",
		})
		w.Write(resp)
		return
	}
	task := model.Task{
		Name:    cr.Name,
		Target:  cr.Target,
		Command: cr.Command,
		Cron:    cr.Cron,
		Runtime: cr.Runtime,
		Status:  constant.TASKSTATUSPENDING,
	}
	err = t.svc.Task().Create(&task)
	if err != nil {
		resp, _ := json.Marshal(types.CommonResp{
			Code:    -1,
			Message: "create task error",
		})
		w.Write(resp)
		return
	}
	resp, _ := json.Marshal(types.CommonResp{
		Code:    0,
		Message: "ok",
	})
	w.Write(resp)
}

// 更新任务
func (t *TaskController) Update(w http.ResponseWriter, req *http.Request) {
	body := make([]byte, req.ContentLength)
	var ur types.UpdateReq
	req.Body.Read(body)
	err := json.Unmarshal(body, &ur)
	if err != nil {
		resp, _ := json.Marshal(types.CommonResp{
			Code:    -1,
			Message: "request error",
		})
		w.Write(resp)
		return
	}
	task, err := t.svc.Task().Get(ur.ID)
	if err != nil {
		resp, _ := json.Marshal(types.CommonResp{
			Code:    -1,
			Message: "get task failed, can't find task by id",
		})
		w.Write(resp)
		return
	}
	task.Status = ur.Status
	task.UpdatedAt = time.Now().Unix()
	err = t.svc.Task().Update(task)
	if err != nil {
		resp, _ := json.Marshal(types.CommonResp{
			Code:    -1,
			Message: "update task failed",
		})
		w.Write(resp)
		return
	}
	resp, _ := json.Marshal(types.CommonResp{
		Code:    0,
		Message: "ok",
	})
	w.Write(resp)
	return
}

// 拉取任务
func (t *TaskController) List(w http.ResponseWriter, req *http.Request) {
	body := make([]byte, req.ContentLength)
	var lr types.ListReq
	req.Body.Read(body)
	err := json.Unmarshal(body, &lr)
	fmt.Println(body)
	if err != nil {
		resp, _ := json.Marshal(types.CommonResp{
			Code:    -1,
			Message: "request error",
		})
		w.Write(resp)
		return
	}
	tasks, err := t.svc.Task().List(lr.Target)
	if err != nil {
		resp, _ := json.Marshal(types.CommonResp{
			Code:    -1,
			Message: "list tasks failed",
		})
		w.Write(resp)
		return
	}
	resp, err := json.Marshal(types.CommonResp{
		Code:    0,
		Message: "ok",
		Data:    tasks,
	})
	w.Write(resp)
	return
}
