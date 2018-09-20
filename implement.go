package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strings"
)

const GlobalCCID = "global"//used in other chaincode, please DONOT remove

var commonLogger = shim.NewLogger("common")

type PayerChainCode struct {
	CommonChaincode
	PayerInterface
	ClinicAuth ClinicAuth
	MemberAuth MemberAuth
	PayerAuth  PayerAuth
}

func (t PayerChainCode) Invoke(stub shim.ChaincodeStubInterface) (response peer.Response) {
	var fcn, params = stub.GetFunctionAndParameters()
	commonLogger.Info("common Invoke:fcn:" + fcn)
	var responseBytes []byte
	switch strings.ToLower(fcn) {
	case "gettokens":
		var tokenVerify, tokenPay = t.GetTokens(t.MemberAuth, params)
		responseBytes = []byte(tokenVerify + "|" + tokenPay)
	case "propose":
		var feeForm = t.Propose(t.ClinicAuth, params)
		responseBytes = []byte(feeForm)
	case "modify":
		var extraFee = t.Modify(t.ClinicAuth, params)
		responseBytes = []byte(extraFee)
	case "revert":
		t.Revert(t.ClinicAuth, params)
	case "settlement":
		t.Settlement(t.PayerAuth, params)
	default:
		PanicString("unknown fcn:" + fcn)
	}
	return shim.Success(responseBytes)
}
