package chaincode

import (
	"strconv"

	"github.com/cdtlab19/kracha/model"
	"github.com/cdtlab19/kracha/router"
	"github.com/cdtlab19/kracha/store"
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
		Handle("CreateCardholder", chaincode.CreateCardholder).
		Handle("GetCardholder", chaincode.GetCardholder).
		Handle("UpdateCardholder", chaincode.UpdateCardholder).
		Handle("DeleteCardholder", chaincode.DeleteCardholder)

	return chaincode
}

func (cc *CardholderChaincode) store(stub shim.ChaincodeStubInterface) *store.CardholderStore {
	return store.NewCardholderStore(stub, cc.logger)
}

//Init implements CardholderChaincode basic setup
func (cc *CardholderChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

//Invoke calls a chaincode method
func (cc *CardholderChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	return cc.router.Invoke(fn, stub, args)
}

//CreateCardholder initiates a basic user defined in model.
// parameters: cpf, name, gender and birthdate
func (cc *CardholderChaincode) CreateCardholder(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 4 {
		cc.logger.Info("Exiting method CreateCardholder.")
		return shim.Error("Incorrect number of arguments; 4 (four) Expected.")
	}

	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" {
		cc.logger.Info("Exiting method CreateCardholder.")
		return shim.Error("Arguments for CreateCardholder can't be null.")
	}

	if _, err := strconv.Atoi(args[0]); err != nil {
		cc.logger.Info("Exiting method CreateCardholder.")
		return shim.Error("Invalid CPF format: must be a numeric string.")
	}

	cpf, name, gender, birthdate := args[0], args[1], args[2], args[3]

	if ch, _ := cc.store(stub).GetCardholder(cpf); ch != nil {
		cc.logger.Info("Exiting method CreateCardholder.")
		return shim.Error("Cardholder already exists.")
	}

	cardholder := model.NewCardholder(cpf, name, gender, birthdate)

	if err := cc.store(stub).SetCardholder(cardholder); err != nil {
		cc.logger.Info("Exiting method CreateCardholder.")
		return shim.Error("Failed to set state for Cardholder: " + err.Error())
	}

	return shim.Success(nil)

}

//GetCardholder returns a cardholder based on it's CPF
// parameters: cpf
func (cc *CardholderChaincode) GetCardholder(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		cc.logger.Info("Exiting method GetCardholder.")
		return shim.Error("Incorrect number of arguments; 1 (one) expected.")
	}

	if args[0] == "" {
		cc.logger.Info("Exiting method GetCardholder.")
		return shim.Error("CPF argument cant be null.")
	}

	if _, err := strconv.Atoi(args[0]); err != nil {
		cc.logger.Info("Exiting method GetCardholder.")
		return shim.Error("Invalid CPF format: must be a numeric string.")
	}

	cpf := args[0]

	cardholder, err := cc.store(stub).GetCardholder(cpf)
	if err != nil {
		cc.logger.Info("Exiting method GetCardholder.")
		return shim.Error("Failed to get cardholder: " + err.Error())
	}

	return shim.Success(cardholder.JSON())

}

//UpdateCardholder updates data from a cardholder based on it's CPF
// parameters: cpf, name, gender and birthdate
func (cc *CardholderChaincode) UpdateCardholder(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		cc.logger.Info("Exiting method UpdateCardholder.")
		return shim.Error("Incorrect number of arguments; 4 (four) Expected.")
	}

	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" {
		cc.logger.Info("Exiting method UpdateCardholder.")
		return shim.Error("Arguments for UpdateCardholder can't be null.")
	}

	if _, err := strconv.Atoi(args[0]); err != nil {
		cc.logger.Info("Exiting method UpdateCardholder.")
		return shim.Error("Invalid CPF format: must be a numeric string.")
	}

	cpf, name, gender, birthdate := args[0], args[1], args[2], args[3]

	if _, err := cc.store(stub).GetCardholder(cpf); err != nil {
		cc.logger.Info("Exiting method UpdateCardholder.")
		return shim.Error("Cardholder doesn't exist.")
	}

	cardholder := model.NewCardholder(cpf, name, gender, birthdate)

	if err := cc.store(stub).SetCardholder(cardholder); err != nil {
		cc.logger.Info("Exiting method UpdateCardholder.")
		return shim.Error("Failed to set state for Cardholder: " + err.Error())
	}

	return shim.Success(nil)
}

//DeleteCardholder deletes a cardholder based on it's CPF
// parameters: cpf
func (cc *CardholderChaincode) DeleteCardholder(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		cc.logger.Info("Exiting method DeleteCardholder")
		return shim.Error("DeleteCardholder expects only 1 argument")
	}

	if _, err := strconv.Atoi(args[0]); err != nil {
		cc.logger.Info("Exiting method DeleteCardholder.")
		return shim.Error("Invalid CPF format: must be a numeric string.")
	}

	cpf := args[0]

	if err := cc.store(stub).DeleteCardholder(cpf); err != nil {
		cc.logger.Info("Exiting method DeleteCardholder.")
		return shim.Error("Coudn't delete Cardholder: " + err.Error())
	}

	return shim.Success(nil)

}
