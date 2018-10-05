package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type PayerChainCode struct {
	*CommonChaincode
	PayerInterface
	ClinicAuth ClinicAuth
	MemberAuth MemberAuth
	PayerAuth  PayerAuth
}

func NewPayerChainCode(name string) PayerChainCode {
	var commonCC = CommonChaincode{}
	commonCC.SetLogger(name)
	return PayerChainCode{CommonChaincode: &commonCC}
}
func (t *PayerChainCode) Prepare(ccAPI shim.ChaincodeStubInterface) {
	t.CommonChaincode.CCAPI = ccAPI
	t.CommonChaincode.Channel = ccAPI.GetChannelID()
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
