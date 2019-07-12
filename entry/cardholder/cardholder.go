package main

import (
	chaincode "github.com/cdtlab19/kracha/chaincode/cardholder"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	logger := shim.NewLogger("cardholder")

	cardholderChaincode := chaincode.NewCardholderChaincode(logger)

	if err := shim.Start(cardholderChaincode); err != nil {
		logger.Critical("[CARDHOLDER] error: %s", err.Error())
	}
}
