package main

import (
    "fmt"
    "rbac-contract/contract"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
    // 创建智能合约
    crct := contract.NewRBACContract()
    // 创建链码
    chaincode, err := contractapi.NewChaincode(crct)
    if err != nil {
        panic(fmt.Sprintf("Error new chaincode: %s\n", err.Error()))
    }
    chaincode.Info.Title = "RBACChaincode"
    chaincode.Info.Version = "0.0.1"
    // 启动链码
    err = chaincode.Start()
    if err != nil {
        panic(fmt.Sprintf("Error starting chaincode: %s\n", err.Error()))
    }
}
