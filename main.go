package main

import (
	"fmt"
	"rbac-contract/rbac"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	// 创建智能合约
	contract := rbac.NewRBACContract()
	// 创建链码
	chaincode, err := contractapi.NewChaincode(contract)
	if err != nil {
		panic(fmt.Sprintf("Error new chaincode: %s\n", err.Error()))
	}
	// 启动链码
	err = chaincode.Start()
	if err != nil {
		panic(fmt.Sprintf("Error starting chaincode: %s\n", err.Error()))
	}
}
