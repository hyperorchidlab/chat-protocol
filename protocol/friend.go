package protocol

type FriendDesc struct {
	PeerPubKey string `json:"peer_pub_key"`
	SendTime   int64  `json:"send_time"`
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
	Op         int        `json:"op"`
	ResultCode int        `json:"result_code"`
	FD         FriendDesc `json:"fd"`
}

type ListFriendReq struct {
	Op int      `json:"op"`
	SP SignPack `json:"sp"`
}

type ListFriendResp struct {
	Op int        `json:"op"`
	SP SignPack   `json:"sp"`
	FL FriendList `json:"fl"`
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
	FD []FriendDetails `json:"fd"`
	GD []GroupDetails  `json:"gd"`
}