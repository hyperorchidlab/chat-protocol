package protocol

import "encoding/json"

const (
	AddFriend int32 = iota
	DelFriend
	AddGroup
	DelGroup
	JoinGroup
	QuitGroup
)


type UserSignText struct {
	AliasName string			`json:"alias_name"`
	CPubKey string				`json:"c_pub_key"`
	SPubKey string				`json:"s_pub_key"`
	ExpireTime int64			`json:"expire_time"`
}

func (ust *UserSignText)ForSig() ([]byte,error)  {
	return json.Marshal(*ust)
}


type SignPack struct {
	Sign string					`json:"sign"`
	SignText UserSignText		`json:"sign_text"`
}


type UserCommand struct {
	Op int				`json:"op"`
	SP SignPack			`json:"sp"`
	CipherTxt []byte 	`json:"cipher_txt"`
}

func NewUserCommand(op int) *UserCommand  {
	uc := &UserCommand{}

	uc.Op = op

	return uc
}

type UCReply struct {
	OP int
	ResultCode int  	`json:"result_code"`
	CipherTxt []byte    `json:"cipher_txt"`
}
