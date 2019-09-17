package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
	. "github.com/davidkhala/goutils"
)

const (
	GlobalCCID      = "global" // used in other chaincode, please DO NOT remove
	FcnPuttoken     = "putToken"
	FcnGettoken     = "getToken"
	FcnRenewtoken   = "renewToken"
	FcnTokenhistory = "tokenHistory"
	FcnDeletetoken  = "deleteToken"
	FcnMovetoken    = "moveToken"
)

func PutToken(t CommonChaincode, token string, tokenData TokenData) {
	var args = ArgsBuilder(FcnPuttoken)
	args.AppendArg(token)
	args.AppendBytes(ToJson(tokenData))
	t.InvokeChaincode(GlobalCCID, args.Get(), "")
}
func RenewToken(t CommonChaincode, token string, newExpiryTime TimeLong) {
	var args = ArgsBuilder(FcnRenewtoken)
	args.AppendArg(token)
	args.AppendArg(newExpiryTime.String())
	t.InvokeChaincode(GlobalCCID, args.Get(), "")
}

func GetToken(t CommonChaincode, token string) *TokenData {
	var args = ArgsBuilder(FcnGettoken)
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
	var args = ArgsBuilder(FcnMovetoken)
	args.AppendArg(token)
	args.AppendBytes(ToJson(request))
	t.InvokeChaincode(GlobalCCID, args.Get(), "")
}
func DeleteToken(t CommonChaincode, token string) {
	var args = ArgsBuilder(FcnDeletetoken)
	args.AppendArg(token)
	t.InvokeChaincode(GlobalCCID, args.Get(), "")
}
func TokenHistory(t CommonChaincode, token string) []KeyModification {
	var args = ArgsBuilder(FcnTokenhistory)
	args.AppendArg(token)

	var payload = t.InvokeChaincode(GlobalCCID, args.Get(), "").Payload
	if payload == nil {
		return nil
	}
	var history []KeyModification
	FromJson(payload, &history)
	return history
}
