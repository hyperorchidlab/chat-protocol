package protocol

import (
	"github.com/kprc/chat-protocol/address"
	"github.com/kprc/chat-protocol/groupid"
)

type GroupMemberDesc struct {
	GroupID  string   `json:"group_id"`
	Friend   string   `json:"friend"`
	SendTime int64    `json:"send_time"`
	Pubkeys  []string `json:"pubkeys"`
	GKeys    []string `json:"g_keys"`
}

type GroupMemberReq struct {
	Op  int             `json:"op"`
	SP  SignPack        `json:"sp"`
	GMD GroupMemberDesc `json:"gmd"`
}

type GroupMbrAddInfo struct {
	GID        groupid.GrpID       `json:"gid"`
	FriendName string              `json:"friend_name"`
	FriendAddr address.ChatAddress `json:"friend_addr"`
	Agree      int                 `json:"agree"`
	JoinTime   int64               `json:"join_time"`
	GKeyHash   string              `json:"g_key_hash"`
}

type GroupMemberResp struct {
	Op         int             `json:"op"`
	ResultCode int             `json:"result_code"`
	GMAI       GroupMbrAddInfo `json:"gmd"`
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
	Owner string    `json:"owner"`
	Gkeys []string  `json:"gkeys"`
	PKeys []string  `json:"pkeys"`
	Hashk string    `json:"hashk"`
	FD    []GMember `json:"fd"`
}

type ListGrpMbrsResp struct {
	Op         int                 `json:"op"`
	ResultCode int                 `json:"result_code"`
	GML        GroupMbrDetailsList `json:"gml"`
}
