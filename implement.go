package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
	. "github.com/davidkhala/goutils"
)

const GlobalCCID = "global" //used in other chaincode, please DONOT remove

type PayerChainCode struct {
	CommonChaincode
	PayerInterface
	ClinicAuth ClinicAuth
	MemberAuth MemberAuth
	PayerAuth  PayerAuth
}

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
func (t PayerAuth) Exec(transient map[string][]byte) bool {
	result := t(transient)
	if ! result {
		PanicString("Payer Authentication failed")
	}
	return result
}
