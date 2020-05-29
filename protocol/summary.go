package protocol

import "github.com/kprc/chat-protocol/groupid"

type FriendSummary struct {
	FriendUpdateTime int64 `json:"fut"`
}

type ListFriendSummaryReq struct {
	Op int      `json:"op"`
	SP SignPack `json:"sp"`
}

type ListFriendSummaryResp struct {
	Op         int           `json:"op"`
	ResultCode int           `json:"result_code"`
	FS         FriendSummary `json:"fs"`
}

type GroupSummary struct {
	GroupId []groupid.GrpID `json:"gids"`
}

type ListGroupSummaryReq struct {
	Op int          `json:"op"`
	SP SignPack     `json:"sp"`
	GS GroupSummary `json:"gs"`
}

type GroupModTime struct {
	GroupId         groupid.GrpID `json:"gid"`
	GroupUpdate     int64         `json:"gu"`
	GrpMemberUpdate int64         `json:"gmu"`
}

type GroupSummaryResult struct {
	GMT []GroupModTime `json:"gmt"`
}

type ListGroupSummaryResp struct {
	Op         int                `json:"op"`
	ResultCode int                `json:"result_code"`
	GSR        GroupSummaryResult `json:"gsr"`
}
