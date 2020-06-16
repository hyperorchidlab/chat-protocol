package protocol

import "github.com/kprc/chat-protocol/groupid"

type GroupMsg struct {
	Gid     groupid.GrpID `json:"gid"`
	AesHash string        `json:"aes"`
	Msg     string        `json:"msg"`
}

type GroupMsgStoreReq struct {
	Op   int      `json:"op"`
	SP   SignPack `json:"sp"`
	GMsg GroupMsg `json:"gmsg"`
}

type GroupMsgStoreResp struct {
	Op         int `json:"op"`
	ResultCode int `json:"result_code"`
}

type GMsgFetch struct {
	Gid   groupid.GrpID `json:"gid"`
	Begin int           `json:"begin"`
	Count int           `json:"count"`
}

type GMsgFetchReq struct {
	Op   int       `json:"op"`
	SP   SignPack  `json:"sp"`
	GMsg GMsgFetch `json:"gmsg"`
}

type LGroupMsg struct {
	AesHash string `json:"aes"`
	Msg     string `json:"msg"`
	Cnt     int    `json:"cnt"`
}

type GMsgFetchContent struct {
	Gid groupid.GrpID `json:"gid"`
	LM  []LGroupMsg   `json:"lm"`
}

type GMsgFetchResp struct {
	Op         int              `json:"op"`
	ResultCode int              `json:"result_code"`
	GMsg       GMsgFetchContent `json:"gmsg"`
}
