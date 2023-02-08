package chaincode

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type Property struct {
	PropertyName string
	AreaSqFt     int
	address      string
	city         string
	state        string
	zipcode      int
	topography   string
	zoning       string
	utilities    string
	access       string
	boundaries   string
}

// Fabric Chaincode interface
type PropertyContract struct {
	contractapi.Contract
	Name               string
	Info               metadata.InfoMetadata
	UnknownTransaction interface{}
	BeforeTransaction  interface{}
	AfterTransaction   interface{}
}

// Init: called when the chaincode is instantiated by the blockchain network
func (c *PropertyContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("Asset modified"))
}

// Invoke
