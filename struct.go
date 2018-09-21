package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
	. "github.com/davidkhala/goutils"
	"strconv"
)

type TokenData struct {
	Owner     string
	OwnerType OwnerType
	TokenType TokenType
	Client    ClientIdentity
}
type OwnerType byte
type TokenType byte

func (t OwnerType) ToString() string {
	switch t {
	case 1:
		return "member"
	case 2:
		return "network"
	case 3:
		return "insurance"
	}
	PanicString("invalid ownerType" + strconv.Itoa(int(t)))
	return ""
}
func (t TokenType) ToString() string {
	switch t {
	case 1:
		return "verify"
	case 2:
		return "pay"
	}
	PanicString("invalid tokenType" + strconv.Itoa(int(t)))
	return ""
}
