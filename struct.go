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

const (
	_ = iota
	OwnerTypeMember
	OwnerTypeNetwork
	OwnerTypeInsurance
)

type TokenType byte

func (t OwnerType) To() string {
	switch t {
	case OwnerTypeMember:
		return "member"
	case OwnerTypeNetwork:
		return "network"
	case OwnerTypeInsurance:
		return "insurance"
	}
	PanicString("invalid ownerType" + strconv.Itoa(int(t)))
	return ""
}
func (OwnerType) New(s string) (OwnerType) {
	var n OwnerType;
	switch s {
	case "member":
		n = OwnerTypeMember
	case "network":
		n = OwnerTypeNetwork
	case "insurance":
		n = OwnerTypeInsurance
	}
	PanicString("invalid ownerType string" + s)
	return n
}

const (
	_ = iota
	TokenTypeVerify
	TokenTypePay
)

func (t TokenType) To() string {
	var s string
	switch t {
	case TokenTypeVerify:
		s = "verify"
	case TokenTypePay:
		s = "pay"
	}
	PanicString("invalid tokenType" + strconv.Itoa(int(t)))
	return s
}
func (TokenType) New(s string) (TokenType) {
	var n TokenType
	switch s {
	case "verify":
		n = TokenTypeVerify
	case "pay":
		n = TokenTypePay
	}
	PanicString("invalid tokenType string:" + s)
	return n
}

type FeeEntry struct {
	Name      string //co-payment | extra-medicine | surgery | diagnose | sick leave days | refer letter
	RawAmount string //filled by clinic, extensible for number handle
	Comment   string //diagnose|refer letter
}
