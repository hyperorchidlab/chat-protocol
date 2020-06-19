package protocol

import "github.com/kprc/chat-protocol/groupid"

type GroupDesc struct {
	GroupAlias string `json:"group_name"`
	GroupID    string `json:"group_id"`
	SendTime   int64  `json:"send_time"`
}

type GroupReq struct {
	Op int       `json:"op"`
	SP SignPack  `json:"sp"`
	GD GroupDesc `json:"gd"`
}

type GroupCreateInfo struct {
	GroupName  string        `json:"group_name"`
	GID        groupid.GrpID `json:"gid"`
	IsOwner    bool          `json:"is_owner"`
	CreateTime int64         `json:"create_time"`
}

type GroupResp struct {
	Op         int             `json:"op"`
	ResultCode int             `json:"result_code"`
	GCI        GroupCreateInfo `json:"gd"`
}
