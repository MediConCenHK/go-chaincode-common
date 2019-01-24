package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang/cid"
	. "github.com/davidkhala/goutils"
)

type TokenData struct {
	Owner      string
	Issuer     string
	Manager    string
	OwnerType  OwnerType
	TokenType  TokenType
	TokenSignature []byte
	ExpiryDate TimeLong
	TransferDate TimeLong
	Client     ClientIdentity
	MetaData   []byte
}

type TokenTransferRequest struct {
	Owner string
	OwnerType
	Manager string
	MetaData   []byte
}

func (t TokenTransferRequest) ApplyOn(data TokenData) TokenData {
	if t.Owner != "" {
		data.Owner = t.Owner
	}
	if t.OwnerType > 0 {
		data.OwnerType = t.OwnerType
	}
	if t.Manager != "" {
		data.Manager = t.Manager
	}
	return data
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
	var enum = []string{"unknown", "member", "clinic", "network", "insurance"}
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
