package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
	. "github.com/davidkhala/goutils"
)

func PutTokenGlobal(t CommonChaincode, token string, tokenData TokenData) {
	var args = ArgsBuilder("put")
	var client = NewClientIdentity(t.CCAPI)
	t.Logger.Info(string(client.ToJSON()))
	tokenData.Client = client
	args.AppendArg("token")
	args.AppendArg(token)
	args.AppendArg(string(ToJson(tokenData)))
	t.InvokeChaincode(GlobalCCID, args.Get(), "") //TODO check response
}

func GetTokenGlobal(t CommonChaincode, token string) (tokenData TokenData) {
	var args = ArgsBuilder("get")

	args.AppendArg("token")
	args.AppendArg(token)
	var payload = t.InvokeChaincode(GlobalCCID, args.Get(), "").Payload
	FromJson(payload, &tokenData)
	return
}
