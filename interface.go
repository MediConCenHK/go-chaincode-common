package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
)

type Payer interface {
	GenTokens(params []string) []byte
}

// NIContract: Network-Payer contract interface
type NIContract interface {
	Propose(params []string) []byte
	Modify(params []string) []byte
	Revert(params []string) []byte
	Settlement(params []string) []byte
}

type VisitData struct {
	Member         string // derived from QRCode,or plain memberData
	Clinic         string // provided by clinic
	Doctor         string // provided by clinic
	MedicalNetwork string // provided by clinic
}

type InsuranceChaincode struct {
	*CommonChaincode
	Payer
}

type NIContractChainCode struct {
	*CommonChaincode
	NIContract
}

func NewNIContract(name string) NIContractChainCode {
	var commonCC = CommonChaincode{}
	commonCC.SetLogger(name)
	return NIContractChainCode{CommonChaincode: &commonCC}
}

func NewInsuranceChaincode(name string) InsuranceChaincode {
	var commonCC = CommonChaincode{}
	commonCC.SetLogger(name)
	return InsuranceChaincode{CommonChaincode: &commonCC}
}
