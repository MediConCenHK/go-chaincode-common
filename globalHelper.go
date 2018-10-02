package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
	. "github.com/davidkhala/goutils"
)

const (
	GlobalCCID        = "global" //used in other chaincode, please DONOT remove
	fcn_putToken      = "putToken"
	fcn_getToken      = "getToken"
	fcn_transferToken = "transferToken"
)

func PutTokenGlobal(t CommonChaincode, token string, tokenData TokenData) {
	var args = ArgsBuilder(fcn_putToken)
	var client = NewClientIdentity(t.CCAPI)
	tokenData.Client = client
	args.AppendArg(token)
	args.AppendBytes(ToJson(tokenData))
	t.InvokeChaincode(GlobalCCID, args.Get(), "") //TODO check response
}

func GetTokenGlobal(t CommonChaincode, token string) (tokenData TokenData) {
	var args = ArgsBuilder(fcn_getToken)
	args.AppendArg(token)
	var payload = t.InvokeChaincode(GlobalCCID, args.Get(), "").Payload
	FromJson(payload, &tokenData)
	return
}
func TransferTokenGlobal(t CommonChaincode, token string, request TokenTransferRequest) (tokenData TokenData) {
	var args = ArgsBuilder(fcn_transferToken)
	args.AppendArg(token)
	args.AppendBytes(ToJson(request))
	var payload = t.InvokeChaincode(GlobalCCID, args.Get(), "").Payload
	FromJson(payload, &tokenData)
	return
}
