package go_chaincode_common

type PayerInterface interface {
	//TODO: is callback designed needed?
	GetTokens(auth MemberAuth, params []string) (tokenVerify, tokenPay string)
	Propose(auth ClinicAuth, params []string) (feeForm string)
	Modify(auth ClinicAuth, params []string) (extraFee string)
	Revert(auth ClinicAuth, params []string)
	Settlement(auth PayerAuth,params []string)
}
type ClinicAuth func(...string) bool
type MemberAuth func(...string) bool
type PayerAuth func(...string) bool

type VisitData struct {
	TokenVerify    string //provided by QRCode
	TokenPay       string //provided by QRCode
	ClinicID       string //provided by clinic
	Doctor         Doctor //provided by clinic
	MedicalNetwork string //provided by clinic
}
type Doctor string //doctor code
