package protocol

import "github.com/kprc/chat-protocol/address"

type FriendDesc struct {
	PeerPubKey string `json:"peer_pub_key"`
	SendTime   int64  `json:"send_time"`
}

type FriendAddInfo struct {
	Addr      address.ChatAddress `json:"addr"`
	AliasName string              `json:"alias_name"`
	AddTime   int64               `json:"add_time"`
	Agree     int                 `json:"agree"`
}

type FriendReq struct {
	Op int        `json:"op"`
	SP SignPack   `json:"sp"`
	FD FriendDesc `json:"fd"`
}

const (
	UserExpired int = iota + 1
	PeerExpired
	PeerNotExisted
	UserNotFound
	InternalError
)

type FriendAddResp struct {
	Op         int           `json:"op"`
	ResultCode int           `json:"result_code"`
	FAI        FriendAddInfo `json:"fd"`
}

type ListFriendReq struct {
	Op int      `json:"op"`
	SP SignPack `json:"sp"`
}

type FriendDetails struct {
	Alias      string `json:"alias"`
	PubKey     string `json:"pub_key"`
	ExpireTime int64  `json:"expire_time"`
	AddTime    int64  `json:"add_time"`
	Agree      int    `json:"agree"` //1 true and false, 0 not sure ,2 false and true, 3 true true
}

type GroupDetails struct {
	Alias      string `json:"alias"`
	IsOwner    bool   `json:"is_owner"`
	CreateTime int64  `json:"create_time"`
	GrpId      string `json:"grp_id"`
	MembrsCnt  int    `json:"membrs_cnt"`
}

type FriendList struct {
	UpdateTime int64           `json:"update_time"`
	FD         []FriendDetails `json:"fd"`
	GD         []GroupDetails  `json:"gd"`
}

type ListFriendResp struct {
	Op int        `json:"op"`
	SP SignPack   `json:"sp"`
	FL FriendList `json:"fl"`
}
