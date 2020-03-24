package rbac

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RBACContract struct {
	contractapi.Contract
}

// 创建一个用于应用链码的合约
func NewRBACContract() *RBACContract {
	contract := &RBACContract{}

	return contract
}

/*
	Init : 初始化合约状态
*/
func (crt *RBACContract) Init() {
	fmt.Println("Running init")
}
