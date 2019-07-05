package chaincode

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type CardholderChaincode struct {
}

func (c *CardholderChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

}
