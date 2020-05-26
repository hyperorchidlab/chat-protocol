package address

import (
	"crypto/ed25519"
	"github.com/btcsuite/btcutil/base58"
)

const (
	AddrPrefix = "BC"
	AddrIDLen  = 45
)

type ChatAddress string

func (addr ChatAddress) String() string {
	return string(addr)
}

func (addr ChatAddress) ToPubKey() ed25519.PublicKey {
	if len(addr) < len(AddrPrefix) {
		return nil
	}
	b58s := string(addr[len(AddrPrefix):])
	return base58.Decode(b58s)
}

func (addr ChatAddress) IsValid() bool {
	if len(addr) <= AddrIDLen {
		return false
	}
	if addr[:len(AddrPrefix)] != AddrPrefix {
		return false
	}
	if len(addr.ToPubKey()) != ed25519.PublicKeySize {
		return false
	}
	return true
}

func ToAddress(key []byte) ChatAddress {
	return ChatAddress(AddrPrefix + base58.Encode(key))
}
