package go_chaincode_common

import (
	"github.com/davidkhala/fabric-common-chaincode-golang"
)

type PayerInterface interface {
	//TODO: is callback designed needed?
	TokenGetter(auth MemberAuth) (tokenVerify, tokenPay string)
	Proposal(auth ClinicAuth) (feeForm string)
	Modifier(auth ClinicAuth, filledFeeForm string) (extraFee string)
	TokenRefund(auth ClinicAuth)
	Settlement(auth PayerAuth)
}
type ClinicAuth func(...string) bool
type MemberAuth func(...string) bool
type PayerAuth func(...string) bool
type PayerChainCode struct {
	golang.CommonChaincode
	PayerInterface
	ClinicAuth ClinicAuth
	MemberAuth MemberAuth
	PayerAuth  PayerAuth
}
type VisitData struct {
	TokenVerify    string //provided by QRCode
	TokenPay       string //provided by QRCode
	ClinicID       string //provided by clinic
	Doctor         Doctor //provided by clinic
	MedicalNetwork string //provided by clinic
}
type Doctor string //doctor code
