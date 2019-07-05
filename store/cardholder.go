package store

import (
	"encoding/json"

	"github.com/cdtlab19/kracha/model"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//CardholderStore abstracts Cardholder CRUD methods
type CardholderStore struct {
	stub   shim.ChaincodeStubInterface
	logger *shim.ChaincodeLogger
}

//NewCardholderStore returns a instance of a Cardholter store
func NewCardholderStore(stub shim.ChaincodeStubInterface, logger *shim.ChaincodeLogger) *CardholderStore {
	return &CardholderStore{stub, logger}
}

func (c *CardholderStore) generateCardholderKey(id string) (key string) {
	key, _ = c.stub.CreateCompositeKey(model.CardholderPrefix, []string{id})
	return
}

//SetCardholder creates or updates a instance of a Cardholders
func (c *CardholderStore) SetCardholder(cardholder *model.Cardholder) error {
	c.logger.Debug("SetCardholder: setting cardholder %s", cardholder.CPF)
	return c.stub.PutState(c.generateCardholderKey(cardholder.CPF), cardholder.JSON())
}

//DeleteCardholder deletes an instance of a Cardholder based on it's CPF
func (c *CardholderStore) DeleteCardholder(CPF string) error {
	c.logger.Debug("DeleteCardholder: deleting cardholder %s", CPF)
	return c.stub.DelState(c.generateCardholderKey(CPF))
}

//GetCardholder returns an instance of a cardholder based on it's CPF
func (c *CardholderStore) GetCardholder(CPF string) (cardholder *model.Cardholder, err error) {
	c.logger.Debug("GetCardholder: getting cardholder %s", CPF)

	data, err := c.stub.GetState(c.generateCardholderKey(CPF))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &cardholder)
	return
}
