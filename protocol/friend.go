package protocol

type FriendDesc struct {
	PeerPubKey string	`json:"peer_pub_key"`
	SendTime int64		`json:"send_time"`
}

type FriendReq struct {
	Op int				`json:"op"`
	SP SignPack			`json:"sp"`
	FD FriendDesc			`json:"fd"`
}

const (
	UserExpired int = iota + 1
	PeerExpired
	PeerNotExisted
	UserNotFound
	InternalError
	)

type FriendAddResp struct {
	Op int				`json:"op"`
	ResultCode int		`json:"result_code"`
	FD FriendDesc		`json:"fd"`
}




