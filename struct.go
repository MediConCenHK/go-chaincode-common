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
type TokenTransferRequest struct {
	FromOwner     string
	FromOwnerType OwnerType
	ToOwner       string
	ToOwnerType   OwnerType
}

type OwnerType byte
type TokenType byte

func (t OwnerType) To() string {
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
func (OwnerType) From(s string) (OwnerType) {
	var n OwnerType;
	switch s {
	case "member":
		n = 1
	case "network":
		n = 2
	case "insurance":
		n = 3
	}
	PanicString("invalid ownerType string" + s)
	return n
}
func (t TokenType) To() string {
	var s string
	switch t {
	case 1:
		s = "verify"
	case 2:
		s = "pay"
	}
	PanicString("invalid tokenType" + strconv.Itoa(int(t)))
	return s
}
func (TokenType) From(s string) (TokenType) {
	var n TokenType
	switch s {
	case "verify":
		n = 1
	case "pay":
		n = 2
	}
	PanicString("invalid tokenType string:" + s)
	return n
}

type FeeEntry struct {
	Name      string //co-payment | extra-medicine | surgery | diagnose | sick leave days | refer letter
	RawAmount string //filled by clinic, extensible for number handle
	Comment   string //diagnose|refer letter
}
