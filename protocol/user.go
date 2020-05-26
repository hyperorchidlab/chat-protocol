package protocol

type UserRegReq struct {
	AliasName    string `json:"alias_name"`
	CPubKey      string `json:"c_pub_key"`
	TimeInterval int64  `json:"time_interval"`
}

type UserRegResp struct {
	SP      SignPack `json:"sp"`
	ErrCode int      `json:"err_code"`
}
