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

type FeeEntry struct {
	Name      string //co-payment | extra-medicine | surgery | diagnose | sick leave days | refer letter
	RawAmount string //filled by clinic, extensible for number handle
	Comment   string //diagnose|refer letter
}
