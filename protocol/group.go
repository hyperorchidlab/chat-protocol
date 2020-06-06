package protocol

type GroupDesc struct {
	GroupAlias string `json:"group_name"`
	GroupID    string `json:"group_id"`
	SendTime   int64 `json:"send_time"`
}

type GroupReq struct {
	Op int       `json:"op"`
	SP SignPack  `json:"sp"`
	GD GroupDesc `json:"gd"`
}

type GroupResp struct {
	Op         int       `json:"op"`
	ResultCode int       `json:"result_code"`
	GD         GroupDesc `json:"gd"`
}
