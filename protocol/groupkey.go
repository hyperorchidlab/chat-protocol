package protocol

type GroupKeys struct {
	GroupKeys []string `json:"group_keys"`
	PubKeys   []string `json:"pub_keys"`
}

type GroupKeyStoreReq struct {
	Op  int       `json:"op"`
	SP  SignPack  `json:"sp"`
	GKs GroupKeys `json:"gks"`
}

type GroupKeyIndex struct {
	IndexKey string `json:"index_key"`
}

type GroupKeyStoreResp struct {
	Op         int           `json:"op"`
	ResultCode int           `json:"result_code"`
	GKI        GroupKeyIndex `json:"gki"`
}

type GroupKeyFetchReq struct {
	Op  int           `json:"op"`
	SP  SignPack      `json:"sp"`
	GKI GroupKeyIndex `json:"gki"`
}

type GroupKeyFetchResp struct {
	Op  int       `json:"op"`
	SP  SignPack  `json:"sp"`
	GKs GroupKeys `json:"gks"`
}
