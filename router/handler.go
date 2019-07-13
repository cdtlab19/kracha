package router

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//Handler defines a Hyperledger Fabric basic type function
type Handler func(stub shim.ChaincodeStubInterface, args []string) peer.Response

//Router handles all Chaincode methods via Invoke
type Router struct {
	routes map[string]Handler
}

//NewRouter creates a basic Chaincode routes
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]Handler),
	}
}

//Handle adds a route to the Router.
//Routes can be overwrited in case on name conflicting
func (r *Router) Handle(method string, handler Handler) *Router {
	r.routes[method] = handler
	return r
}

//Invoke returns a Chaincode method based on it's initial name
func (r *Router) Invoke(method string, stub shim.ChaincodeStubInterface, args []string) peer.Response {
	f, ok := r.routes[method]
	if !ok {
		//throw some error
	}

	return f(stub, args)
}
