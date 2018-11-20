package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
	. "github.com/davidkhala/goutils"
)

type TokenData struct {
	Owner      string
	Issuer     string
	Manager    string
	OwnerType  OwnerType
	TokenType  TokenType
	ExpireDate TimeLong
	Client     ClientIdentity
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
	OwnerTypeClinic
	OwnerTypeNetwork
	OwnerTypeInsurance
)

type TokenType byte

func (t OwnerType) To() string {
	var enum = []string{"member", "network", "clinic", "insurance"}
	return enum[t]
}

const (
	_ = iota
	TokenTypeVerify
	TokenTypePay
)

func (t TokenType) To() string {
	var enum = []string{"verify", "pay"}
	return enum[t]
}
func (TokenType) From(s string) TokenType {
	var typeMap = map[string]TokenType{"verify": TokenTypeVerify, "pay": TokenTypePay}
	return typeMap[s]
}

type FeeEntry struct {
	Name      string //co-payment | extra-medicine | surgery | diagnose | sick leave days | refer letter
	RawAmount string //filled by clinic, extensible for number handle
	Comment   string //diagnose|refer letter
}
