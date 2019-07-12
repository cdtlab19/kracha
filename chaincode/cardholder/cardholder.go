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
	chaincode.router = router.NewRouter().
		Handle("CreateUser", chaincode.CreateUser).
		Handle("GetUser", chaincode.GetUser)

	return chaincode
}

//Init implements CardholderChaincode basic setup
func (c *CardholderChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

//Invoke calls a chaincode method
func (c *CardholderChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, _ := stub.GetFunctionAndParameters()
	return c.router.Invoke(fn, stub)
}

// CreateUser initiates a basic user defined in model
func (c *CardholderChaincode) CreateUser(stub shim.ChaincodeStubInterface) peer.Response {

}

// GetUser returns a user based on it's CPF
func (c *CardholderChaincode) GetUser(stub shim.ChaincodeStubInterface) peer.Response {
}
