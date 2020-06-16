package protocol

type P2pMsg struct {
	PeerPk string `json:"ppk"`
	MyPk   string `json:"mpk"`
	Msg    string `json:"msg"`
}

type P2pMsgStoreReq struct {
	Op  int      `json:"op"`
	SP  SignPack `json:"sp"`
	Msg P2pMsg   `json:"msg"`
}

type P2pMsgStoreResp struct {
	Op         int `json:"op"`
	ResultCode int `json:"result_code"`
}

type P2pMsgFetch struct {
	Begin int `json:"begin"`
	Count int `json:"count"`
}

type P2pMsgFetchReq struct {
	Op  int         `json:"op"`
	SP  SignPack    `json:"sp"`
	Msg P2pMsgFetch `json:"msg"`
}

type LP2pMsg struct {
	PeerPk string `json:"ppk"`
	Msg    string `json:"msg"`
	Cnt    int    `json:"cnt"`
}

type P2pMsgFetchResp struct {
	Op         int       `json:"op"`
	ResultCode int       `json:"result_code"`
	Msg        []LP2pMsg `json:"msg"`
}
