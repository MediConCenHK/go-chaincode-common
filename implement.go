package go_chaincode_common

import (
	. "github.com/davidkhala/fabric-common-chaincode-golang"
)

type InsuranceChaincode struct {
	*CommonChaincode
	MemberAuth
	InsuranceAuth
}
type NetworkChainCode struct {
	*CommonChaincode
	ClinicAuth
	NetworkAuth
}
type NIContractChainCode struct {
	*CommonChaincode
	NIContract
	NetworkAuth
	InsuranceAuth
}

func NewNIContract(name string) NIContractChainCode {
	var commonCC = CommonChaincode{}
	commonCC.SetLogger(name)
	return NIContractChainCode{CommonChaincode: &commonCC}
}
func NewNetworkChainCode(name string) NetworkChainCode {
	var commonCC = CommonChaincode{}
	commonCC.SetLogger(name)
	return NetworkChainCode{CommonChaincode: &commonCC}
}
func NewInsuranceChaincode(name string) InsuranceChaincode {
	var commonCC = CommonChaincode{}
	commonCC.SetLogger(name)
	return InsuranceChaincode{CommonChaincode: &commonCC}
}
