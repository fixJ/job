package types

type NodeInfo struct {
	IP        string
	CreatedAt int64
	LastProbe int64
}

type CommonResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type CreateReq struct {
	Name    string `json:"name"`
	Target  string `json:"target"`
	Command string `json:"command"`
	Cron    bool   `json:"cron"`
	Runtime int64  `json:"runtime"`
}

type UpdateReq struct {
	ID     int64 `json:"id"`
	Status int   `json:"status"`
}

type ListReq struct {
	Target string `json:"target"`
}

type LiveReq struct {
	Target string `json:"target"`
}
