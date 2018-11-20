package go_chaincode_common

import (
	. "github.com/davidkhala/goutils"
)

type Payer interface {
	GenTokens(auth MemberAuth, params []string) []byte
	GetMemberData(params []string) []byte
}

//NIContract: Network-Insurance contract interface
type NIContract interface {
	Propose(auth ClinicAuth, params []string) []byte
	Modify(auth ClinicAuth, params []string) []byte
	Revert(auth ClinicAuth, params []string) []byte
	Settlement(auth PayerAuth, params []string) []byte
}

const (
	Payer_fcn_genTokens     = "genTokens"
	Payer_fcn_getMemberData = "getMemberData"
	Contract_fcn_propose    = "propose"
	Contract_fcn_modify     = "modify"
	Contract_fcn_revert     = "revert"
	Contract_fcn_settlement = "settlement"
)

type ClinicAuth func(transient map[string][]byte) bool
type MemberAuth func(transient map[string][]byte) bool
type NetworkAuth func(transient map[string][]byte) bool
type InsuranceAuth func(transient map[string][]byte) bool
type PayerAuth func(transient map[string][]byte) bool //TODO for settlement extension

func (t ClinicAuth) Exec(transient map[string][]byte) bool {
	result := t(transient)
	if ! result {
		PanicString("Clinic Authentication failed")
	}
	return result
}
func (t MemberAuth) Exec(transient map[string][]byte) bool {
	result := t(transient)
	if ! result {
		PanicString("Member Authentication failed")
	}
	return result
}
func (t NetworkAuth) Exec(transient map[string][]byte) bool {
	result := t(transient)
	if ! result {
		PanicString("Network Authentication failed")
	}
	return result
}
func (t InsuranceAuth) Exec(transient map[string][]byte) bool {
	result := t(transient)
	if ! result {
		PanicString("Insurance Authentication failed")
	}
	return result
}

type VisitData struct {
	Member         string //derived from QRCode,or plain memberData
	Clinic         string //provided by clinic
	Doctor         string //provided by clinic
	MedicalNetwork string //provided by clinic
}
