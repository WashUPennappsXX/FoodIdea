/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

type food producer struct {
	FoodCategory   string  `json:"docType"`
	FoodName    string  `json:"name"`
	Distance     string  `json:"distant"`
	
}
//需求方
type food delivery struct {
	FoodCategory   string  `json:"docType"`
	FoodName    string  `json:"name"`
	Review        string 'json.review'
	Rating        double    `json:"rating"`
	Distance     string  `json:"distant"`
	Price  		int     `json:"price"`

	
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fcn, args := stub.GetFunctionAndParameters()

	if fcn == "initorder" {
		return cc.initorder(stub, args)
	} 
	else{
	return shim.Error("Invoke fcn error")
}

func (cc *Chaincode) initorder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 9 {
		return shim.Error("not enough args")
	}
	FoodCategory   := strconv.Atoi(args[0])
	FoodName := strings.ToLower(args[1])
	category := strings.ToLower(args[2])
	Rating := strconv.Atoi(args[3])
	distance := strings.ToLower(args[4])
	price := strconv.Atoi(args[5])


	// ==== Create demond object and marshal to JSON ====
	objectType := "order"
	order := &order{objectType, FoodCategory, FoodName, category, rating, distance, price}
	orderJSONasBytes, err := json.Marshal(order)
	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save order to state ===
	err = stub.PutState(strconv.Itoa(orderID), orderJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init order")
	return shim.Success(nil)
}



	
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	//value1 := []byte{0x00}
	stub.PutState(nameUnpsupplyIDIndexKey, unprofitablesupplyJSONasBytes)

	indexName2 := "foodproducer~unpsupplyID"
	foodproducerUnpsupplyIDIndexKey, err := stub.CreateCompositeKey(indexName2, []string{strconv.Itoa(unprofitablesupply.ProviderID), strconv.Itoa(unprofitablesupply.SupplyID)})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value

	stub.PutState(foodproducerUnpsupplyIDIndexKey, unprofitablesupplyJSONasBytes)
	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init unprofitablesupply")
	return shim.Success(nil)
}

//创建供应方
func (cc *Chaincode) initDelivery(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 7 {
		return shim.Error("not enough args")
	}
	FoodCategory   := strconv.Atoi(args[0])
	FoodName := strings.ToLower(args[1])
	category := strings.ToLower(args[2])
	Rating := strconv.Atoi(args[3])
	distance := strings.ToLower(args[4])
	price := strconv.Atoi(args[5])
	
	
	// ==== Create demand object and marshal to JSON ====
	objectType := "foodproducer"
	foodproducer := &foodproducer{FoodCategory, FoodName, distance}
	foodproducerJSONasBytes, err := json.Marshal(foodproducer)
	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save order to state ===
	err = stub.PutState(strconv.Itoa(orgID), foodproducerJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init foodproducer")
	return shim.Success(nil)
}

func (cc *Chaincode) initProvider(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 8 {
		return shim.Error("not enough args")
	}
	FoodCategory, _ := strconv.Atoi(args[0])
	FoodName = strconv.Atoi(args[0])
	name := strings.ToLower(args[1])
	
	Distance, _ := strconv.ParseFloat(args[7], 64)


	providerAsBytes, err := stub.GetState(strconv.Itoa(orgID))
	if err != nil {
		return shim.Error("Failed to get provider: " + err.Error())
	} else if providerAsBytes != nil {
		fmt.Println("This provider already exists: " + strconv.Itoa(orgID))
		return shim.Error("This provider already exists: " + strconv.Itoa(orgID))
	}



	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init foodproducer")
	return shim.Success(nil)
}

func (cc *Chaincode) queryBycategory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var User string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error(err.Error())
	}

	User = args[0]

	// Get the state from the ledger
	value, err := stub.GetState(User)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println(string(value))
	return shim.Success([]byte(value))
}

func (cc *Chaincode) queryByName(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	name := strings.ToLower(args[0])

	fmt.Println("- start queryUnproByName ", name)

	// Query the color~name index by color
	// This will execute a key range query on all keys starting with 'color'
	nameUnpsupplyIDResultsIterator, err := stub.GetStateByPartialCompositeKey("name~unpsupplyID", []string{name})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer nameUnpsupplyIDResultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(nameUnpsupplyIDResultsIterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return shim.Success([]byte(buffer.Bytes()))
}


func (cc *Chaincode) queryByfoodproducer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	foodproducer := strings.ToLower(args[0])

	fmt.Println("- start queryProBqueryProByfoodproduceryName ", foodproducer)

	// Query the color~name index by color
	// This will execute a key range query on all keys starting with 'color'
	foodproducerPsupplyIDResultsIterator, err := stub.GetStateByPartialCompositeKey("foodproducer~unpsupplyID", []string{foodproducer})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer foodproducerPsupplyIDResultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(foodproducerPsupplyIDResultsIterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return shim.Success([]byte(buffer.Bytes()))
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, supress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

/
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

//存储数据，2个参数  
func (cc *Chaincode) set(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting key and value")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

//删除信息，一个参数
func (cc *Chaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.Expecting the key to delete")
	}

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}
