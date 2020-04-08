package contract

import api "github.com/hyperledger/fabric-contract-api-go/contractapi"

type RoleListInterface interface {
    AddRole(role *Role) error
    GetRole(roleId string) (*Role, error)
    DelRole(roleId string) error
}

type RoleList struct {
    ctx api.TransactionContextInterface
}

func newRoleList(ctx api.TransactionContextInterface) *RoleList {
    lst := &RoleList{ctx: ctx}
    return lst
}

func (rl *RoleList) AddRole(role *Role) error {
    panic("implement me")
}

func (rl *RoleList) GetRole(roleId string) (*Role, error) {
    panic("implement me")
}

func (rl *RoleList) DelRole(roleId string) error {
    panic("implement me")
}
