package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
)

const GlobalCCID = "global" //used in other chaincode, please DONOT remove

type PayerChainCode struct {
	CommonChaincode
	PayerInterface
	ClinicAuth ClinicAuth
	MemberAuth MemberAuth
	PayerAuth  PayerAuth
}

