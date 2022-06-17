package controller

type CommonResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CreateReq struct {
	Name    string `json:"name"`
	Target  string `json:"target"`
	Command string `json:"command"`
	Cron    bool   `json:"cron"`
	Runtime int64  `json:"runtime"`
}
