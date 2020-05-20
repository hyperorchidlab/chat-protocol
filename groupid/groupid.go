package groupid

import (
	"crypto/rand"
	"github.com/btcsuite/btcutil/base58"
)


const(
	GrpIDPrefix = "gc"
	GrpIDLen = 45
	GrpIDBytesLen = 32
)

type GrpID string

func NewGroupId() GrpID  {
	buf := make([]byte,GrpIDBytesLen)
	for{
		n,_:=rand.Read(buf)
		if n != len(buf){
			continue
		}else {
			break
		}
	}

	return ToGroupID(buf)

}

func (gid GrpID)ToBytes() []byte  {
	return base58.Decode(gid.String()[len(GrpIDPrefix):])
}


func (gid GrpID) IsValid() bool {
	if len(gid) < GrpIDLen {
		return false
	}
	if gid[:len(GrpIDPrefix)] != GrpIDPrefix {
		return false
	}
	if len(gid.ToBytes()) != GrpIDBytesLen {
		return false
	}
	return true
}

func (gid GrpID)String() string  {
	return string(gid)
}

func ToGroupID(buf []byte) GrpID  {
	return GrpID(GrpIDPrefix+base58.Encode(buf))
}



