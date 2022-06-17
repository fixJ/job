package controller

import (
	"encoding/json"
	"job/internal/jobserver/manager"
	"net/http"
)

type CoreController struct {
}

func NewCoreController() *CoreController {
	return &CoreController{}
}

// 节点注册
func (c *CoreController) LiveProbe(w http.ResponseWriter, req *http.Request) {
	body := make([]byte, req.ContentLength)
	var lr LiveReq
	req.Body.Read(body)
	err := json.Unmarshal(body, &lr)
	if err != nil {
		resp, _ := json.Marshal(CommonResp{
			Code:    -1,
			Message: "request error",
		})
		w.Write(resp)
		return
	}
	m, err := manager.GetManagerOr()
	if err != nil {
		resp, _ := json.Marshal(CommonResp{
			Code:    -1,
			Message: "get manager failed",
		})
		w.Write(resp)
		return
	}
	m.UpdateLiveNode(lr.Target)
	resp, _ := json.Marshal(CommonResp{
		Code:    0,
		Message: "ok",
	})
	w.Write(resp)
}