package go_chaincode_common

import (
	"github.com/davidkhala/fabric-common-chaincode-golang"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type PayerChainCode struct {
	golang.CommonChaincode
	PayerInterface
	ClinicAuth ClinicAuth
	MemberAuth MemberAuth
	PayerAuth  PayerAuth
	Name       string
	Logger     *shim.ChaincodeLogger
}
