package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
	. "github.com/davidkhala/goutils"
)

// used in other chaincode, please DO NOT remove
const (
	GlobalCCID      = "global"
	FcnCreateToken  = "createToken"
	FcnGetToken     = "getToken"
	FcnRenewToken   = "renewToken"
	FcnTokenHistory = "tokenHistory"
	FcnDeleteToken  = "deleteToken"
	FcnMoveToken    = "moveToken"
)

func CreateToken(t CommonChaincode, token string, request TokenCreateRequest) {
	var args = ArgsBuilder(FcnCreateToken)
	args.AppendArg(token)
	args.AppendBytes(ToJson(request))
	t.InvokeChaincode(GlobalCCID, args.Get(), "")
}
func RenewToken(t CommonChaincode, token string, newExpiryTime TimeLong) {
	var args = ArgsBuilder(FcnRenewToken)
	args.AppendArg(token)
	args.AppendArg(newExpiryTime.String())
	t.InvokeChaincode(GlobalCCID, args.Get(), "")
}

func GetToken(t CommonChaincode, token string) *TokenData {
	var args = ArgsBuilder(FcnGetToken)
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
	var args = ArgsBuilder(FcnMoveToken)
	args.AppendArg(token)
	args.AppendBytes(ToJson(request))
	t.InvokeChaincode(GlobalCCID, args.Get(), "")
}
func DeleteToken(t CommonChaincode, token string) {
	var args = ArgsBuilder(FcnDeleteToken)
	args.AppendArg(token)
	t.InvokeChaincode(GlobalCCID, args.Get(), "")
}
func TokenHistory(t CommonChaincode, token string) []KeyModification {
	var args = ArgsBuilder(FcnTokenHistory)
	args.AppendArg(token)

	var payload = t.InvokeChaincode(GlobalCCID, args.Get(), "").Payload
	if payload == nil {
		return nil
	}
	var history []KeyModification
	FromJson(payload, &history)
	return history
}
