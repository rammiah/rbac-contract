package rbac

import (
	"fmt"

	api "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RBACContract struct {
	api.Contract
}

// 创建一个用于应用链码的合约
func NewRBACContract() *RBACContract {
	contract := &RBACContract{}
	contract.TransactionContextHandler = &RBACContract{}

	return contract
}


// 初始化合约状态
func (crt *RBACContract) Init(ctx RBACContext) {
	fmt.Println("Running init")
}

// 添加单个用户
func (crt *RBACContract) AddUser() {

}

// 批量添加用户
func (crt *RBACContract) AddUsers() {

}
