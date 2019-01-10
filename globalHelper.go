package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
	. "github.com/davidkhala/goutils"
)

const (
	GlobalCCID       = "global" //used in other chaincode, please DONOT remove
	Fcn_putToken     = "putToken"
	Fcn_getToken     = "getToken"
	Fcn_renewToken   = "renewToken"
	Fcn_tokenHistory = "tokenHistory"
	Fcn_deleteToken  = "deleteToken"
	Fcn_moveToken    = "moveToken"
)

func PutToken(t CommonChaincode, token string, tokenData TokenData) {
	var args = ArgsBuilder(Fcn_putToken)
	args.AppendArg(token)
	args.AppendBytes(ToJson(tokenData))
	t.InvokeChaincode(GlobalCCID, args.Get(), "")
}
func RenewToken(t CommonChaincode, token string, newExpiryTime TimeLong) {
	var args = ArgsBuilder(Fcn_renewToken)
	args.AppendArg(token)
	args.AppendArg(newExpiryTime.String())
	t.InvokeChaincode(GlobalCCID, args.Get(), "")
}

func GetToken(t CommonChaincode, token string) (*TokenData) {
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
func MoveToken(t CommonChaincode, token string, request TokenTransferRequest) {
	var args = ArgsBuilder(Fcn_moveToken)
	args.AppendArg(token)
	args.AppendBytes(ToJson(request))
	t.InvokeChaincode(GlobalCCID, args.Get(), "") //TODO check response
}
func DeleteToken(t CommonChaincode, token string) {
	var args = ArgsBuilder(Fcn_deleteToken)
	args.AppendArg(token)
	t.InvokeChaincode(GlobalCCID, args.Get(), "")
}
func TokenHistory(t CommonChaincode, token string) []KeyModification {
	var args = ArgsBuilder(Fcn_tokenHistory)
	args.AppendArg(token)

	var payload = t.InvokeChaincode(GlobalCCID, args.Get(), "").Payload
	if payload == nil {
		return nil
	}
	var history []KeyModification
	FromJson(payload, &history)
	return history
}
