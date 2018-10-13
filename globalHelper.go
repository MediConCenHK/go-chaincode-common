package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
	. "github.com/davidkhala/goutils"
)

const (
	GlobalCCID        = "global" //used in other chaincode, please DONOT remove
	Fcn_putToken      = "putToken"
	Fcn_getToken      = "getToken"
	Fcn_transferToken = "transferToken"
)

func PutTokenGlobal(t CommonChaincode, token string, tokenData *TokenData) {
	if token == "" || tokenData == nil {
		return
	}
	var args = ArgsBuilder(Fcn_putToken)
	var client = NewClientIdentity(t.CCAPI)
	tokenData.Client = client
	args.AppendArg(token)
	args.AppendBytes(ToJson(*tokenData))
	t.InvokeChaincode(GlobalCCID, args.Get(), "") //TODO check response
}

func GetTokenGlobal(t CommonChaincode, token string) (*TokenData) {
	if token == "" {
		return nil
	}
	var args = ArgsBuilder(Fcn_getToken)
	args.AppendArg(token)
	var payload = t.InvokeChaincode(GlobalCCID, args.Get(), "").Payload
	if payload == nil {
		return nil
	}
	var tokenData TokenData
	FromJson(payload, &tokenData)
	return &tokenData
}
func TransferTokenGlobal(t CommonChaincode, token string, request TokenTransferRequest) (*TokenData) {
	var args = ArgsBuilder(Fcn_transferToken)
	args.AppendArg(token)
	args.AppendBytes(ToJson(request))
	var payload = t.InvokeChaincode(GlobalCCID, args.Get(), "").Payload
	if payload == nil {
		return nil
	}
	var tokenData TokenData
	FromJson(payload, &tokenData)
	return &tokenData
}
