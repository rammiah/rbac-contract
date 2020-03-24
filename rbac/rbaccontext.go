package rbac

import (
	api "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RBACContextInterface interface {
	api.TransactionContextInterface
	AddUser() error
	AddUsers() error
}

// rbac模型的上下文
type RBACContext struct {
	api.TransactionContext
	// 加入用户，角色，权限，文件等资源的数据
	users []User
}