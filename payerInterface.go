package go_chaincode_common

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)
type TokenGetter func() string
type Proposal func() string

type InsuranceChaincode interface {
	shim.Chaincode
}
