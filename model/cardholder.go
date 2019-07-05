package model

import "encoding/json"

//CardholderPrefix avoids double meaning
const CardholderPrefix = "USER"

//Cardholder defines a credit card owner
type Cardholder struct {
	DocType   string `json:"docType"`
	CPF       string `json:"cpf"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
}

//NewCardholder returns a Cardholder instance
func NewCardholder(cpf string, name string, gender string, birthdate string) *Cardholder {
	return &Cardholder{
		DocType:   CardholderPrefix,
		CPF:       cpf,
		Name:      name,
		Gender:    gender,
		Birthdate: birthdate,
	}
}

//JSON encodes a Cardholder model as a JSON object
func (c *Cardholder) JSON() []byte {
	v, _ := json.Marshal(c)
	return v
}
