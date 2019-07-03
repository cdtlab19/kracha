package model

import "encoding/json"

//UserPrefix avoids double meaning
const UserPrefix = "USER"

//Cardholder defines a credit card owner
type Cardholder struct {
	Type      string `json:"type"`
	CPF       string `json:"cpf"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
}

//NewCardholder returns a Cardholder instance
func NewCardholder(cpf string, name string, birthdate string) *Cardholder {
	return &Cardholder{
		Type:      UserPrefix,
		CPF:       cpf,
		Name:      name,
		Birthdate: birthdate,
	}
}

//JSON encodes a Cardholder model as a JSON object
func (c *Cardholder) JSON() []byte {
	v, _ := json.Marshal(c)
	return v
}
