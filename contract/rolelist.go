package contract

import (
    api "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "encoding/json"
)

type RoleListInterface interface {
    AddRole(role *Role) error
    GetRole(roleId string) (*Role, error)
    DelRole(roleId string) error
}

type RoleList struct {
    ctx api.TransactionContextInterface
    Name string
}

func newRoleList(ctx api.TransactionContextInterface) *RoleList {
    lst := &RoleList{
        ctx: ctx,
        Name: "org.rammiah.rolelist",
    }
    return lst
}

func (rl *RoleList) AddRole(role *Role) error {
    key, err := rl.ctx.GetStub().CreateCompositeKey(rl.Name, []string{role.ID})
    if err != nil {
        return err
    }
    data, err := json.Marshal(role)
    if err != nil {
        return err
    }
    return rl.ctx.GetStub().PutState(key, data)
}

func (rl *RoleList) GetRole(id string) (*Role, error) {
    key, err := rl.ctx.GetStub().CreateCompositeKey(rl.Name, []string{id})
    if err != nil {
        return nil, err
    }
    data, err := rl.ctx.GetStub().GetState(key)
    if err != nil {
        return nil, err
    }
    role := new(Role)
    err = json.Unmarshal(data, role)
    return role, err
}

func (rl *RoleList) DelRole(id string) error {
    key, err := rl.ctx.GetStub().CreateCompositeKey(rl.Name, []string{id})
    if err != nil {
        return err
    }
    return rl.ctx.GetStub().DelState(key)
}
