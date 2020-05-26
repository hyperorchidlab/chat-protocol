package protocol

type GroupMemberDesc struct {
	GroupID  string `json:"group_id"`
	Friend   string `json:"friend"`
	SendTime string `json:"send_time"`
}

type GroupMemberReq struct {
	Op  int             `json:"op"`
	SP  SignPack        `json:"sp"`
	GMD GroupMemberDesc `json:"gmd"`
}

type GroupMemberResp struct {
	Op         int             `json:"op"`
	ResultCode int             `json:"result_code"`
	GMD        GroupMemberDesc `json:"gmd"`
}

type ListGrpMbr struct {
	GroupId string `json:"group_id"`
}

type ListGroupMbrsReq struct {
	Op int        `json:"op"`
	SP SignPack   `json:"sp"`
	LG ListGrpMbr `json:"lg"`
}

type GMember struct {
	Alias      string `json:"alias"`
	PubKey     string `json:"pub_key"`
	ExpireTime int64  `json:"expire_time"`
	Agree      int    `json:"agree"` //1 true and false, 0 not sure ,2 false and true, 3 true true
}

type GroupMbrDetailsList struct {
	FD []GMember
}

type ListGrpMbrsResp struct {
	Op         int                 `json:"op"`
	ResultCode int                 `json:"result_code"`
	GML        GroupMbrDetailsList `json:"gml"`
}
