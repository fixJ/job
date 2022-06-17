package types

type TaskInfo struct {
	ID      string `json:"id"`
	Command string `json:"command"`
	Cron    bool   `json:"cron"`
	Runtime int64  `json:"runtime"`
}

type LiveReq struct {
	Target string `json:"target"`
}
