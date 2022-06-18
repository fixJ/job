package types

type TaskInfo struct {
	ID      int64  `json:"id"`
	Command string `json:"command"`
	Cron    bool   `json:"cron"`
	Runtime int64  `json:"runtime"`
}

type LiveReq struct {
	Target string `json:"target"`
}

type ListReq struct {
	Target string `json:"target"`
}

type UpdateReq struct {
	ID     int64 `json:"id"`
	Status int   `json:"status"`
}

type ListResp struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []TaskInfo `json:"data,omitempty"`
}
