package go_chaincode_common

import (
	"errors"
	"fmt"
	. "github.com/davidkhala/fabric-common-chaincode-golang"
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	logger "github.com/sirupsen/logrus"
	"testing"
)

const (
	testName = "test"
)

type TestChaincode struct {
	NIContractChainCode
}

func (t *TestChaincode) Init(stub shim.ChaincodeStubInterface) (response peer.Response) {
	defer Deferred(DeferHandlerPeerResponse, &response)
	t.Prepare(stub)
	logger.Info("Init")
	return shim.Success(nil)
}
func (t *TestChaincode) Invoke(stub shim.ChaincodeStubInterface) (response peer.Response) {
	defer Deferred(DeferHandlerPeerResponse, &response)
	var fcn, _ = stub.GetFunctionAndParameters()
	logger.Info("Invoke fcn: " + fcn)
	t.Prepare(stub)
	var responseBytes []byte
	switch fcn {
	case "panic":
		response = shim.Error("anyerror")
		panic(errors.New("testPanic"))
	default:
		PanicString("unknown fcn:" + fcn)
	}
	return shim.Success(responseBytes)
}

var cc = TestChaincode{
	NewNIContract(),
}

func TestInit(t *testing.T) {
	var mock = shim.NewMockStub(testName, &cc)
	var args = ArgsBuilder("init")
	mock.MockInit("initTxId", args.Get())
}
func TestInvokePanic(t *testing.T) {
	var mock = shim.NewMockStub(testName, &cc)

	var args = ArgsBuilder("panic")
	var response = mock.MockInvoke("invokeTxid", args.Get())
	fmt.Println(response)
}
