package chaincode

import (
	"github.com/cdtlab19/kracha/router"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//CardholderChaincode abstracts chaincode calls
type CardholderChaincode struct {
	logger *shim.ChaincodeLogger
	router *router.Router
}

//NewCardholderChaincode creates a new Chaincode with it's predefined routes
func NewCardholderChaincode(logger *shim.ChaincodeLogger) *CardholderChaincode {
	chaincode := &CardholderChaincode{
		logger: logger,
	}
	chaincode.router.Handle("CreateUser", chaincode.CreateUser).
		Handle("GetUset", chaincode.GetUser)

	return chaincode
}

//Init implements CardholderChaincode basic setup
func (c *CardholderChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

//CreateUser initiates a basic user defined in model
func CreateUser(stub shim.ChaincodeStubInterface) peer.Response {
	//TODO
}

//GetUser returns a user based on it's CPF
func GetUser(stub shim.ChaincodeStubInterface) peer.Response {
	//TODO
}
