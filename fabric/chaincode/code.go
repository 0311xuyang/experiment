package main

import (
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Chaincode struct {
}

func (c *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (c *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "decrease" {
		return c.decrease(stub, args)
	} else if function == "increase" {
		return c.increase(stub, args)
	} else if function == "query" {
		return c.query(stub, args)
	}
	return shim.Error("Invalid invoke function")
}

func (c *Chaincode) decrease(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string
	var Aval int
	if len(args) != 1 {
		return shim.Error("Invalid increase param")
	}
	A = args[0]
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))
	Aval -= 1
	fmt.Printf("Aval = %d\n", Aval)
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		shim.Error(err.Error())
	}
	Avalbytes, _ = stub.GetState(A)
	return shim.Success(Avalbytes)
}

func (c *Chaincode) increase(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string
	var Aval int
	if len(args) != 1 {
		return shim.Error("Invalid increase param")
	}
	A = args[0]
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		Aval = 0
	} else {
		Aval, _ = strconv.Atoi(string(Avalbytes))
	}
	Aval += 1
	fmt.Printf("Aval = %d\n", Aval)
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (c *Chaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string
	var err error
	if len(args) != 1 {
		return shim.Error("Invalid query param")
	}
	A = args[0]
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}
	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}
	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
