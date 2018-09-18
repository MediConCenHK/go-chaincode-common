package go_chaincode_common

import (
	"github.com/davidkhala/fabric-common-chaincode-golang"
)

type PayerInterface interface {
	//TODO: is callback designed needed?
	TokenGetter(auth MemberAuth)
	Proposal(auth ClinicAuth)
	Modifier(auth ClinicAuth)
	TokenRefund(auth ClinicAuth)
	Settlement(auth PayerAuth)
}
type ClinicAuth func(...string) bool
type MemberAuth func(...string) bool
type PayerAuth func(...string) bool
type PayerChainCode struct {
	golang.CommonChaincode
	PayerInterface
}
