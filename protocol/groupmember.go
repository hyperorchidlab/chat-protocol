package protocol


type GroupMemberDesc struct {
	owner string	`json:"owner"`
	GroupID string  `json:"group_id"`
	Friend  string  `json:"friend"`
	SendTime string `json:"send_time"`
}



type GroupMemberReq struct {
	Op int					`json:"op"`
	SP SignPack				`json:"sp"`
	GMD GroupMemberDesc		`json:"gmd"`
}

type GroupMemberResp struct {
	Op int					`json:"op"`
	ResultCode int  		`json:"result_code"`
	GMD GroupMemberDesc		`json:"gmd"`
}
