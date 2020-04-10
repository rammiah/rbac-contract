package contract

import (
    "encoding/json"
    api "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PermissionListInterface interface {
    AddPermission(permission *Permission) error
    GetPermission(pId string) (*Permission, error)
    DelPermission(pId string) error
}

type PermissionList struct {
    ctx  api.TransactionContextInterface
    Name string
}

func newPermissionList(ctx api.TransactionContextInterface) *PermissionList {
    lst := &PermissionList{
        ctx:  ctx,
        Name: "org.rammiah.permissionlist",
    }

    return lst
}

func (pl *PermissionList) AddPermission(permission *Permission) error {
    //panic("implement me")
    key, err := pl.ctx.GetStub().CreateCompositeKey(pl.Name, []string{permission.ID})
    if err != nil {
        return err
    }
    data, err := json.Marshal(permission)
    if err != nil {
        return err
    }
    return pl.ctx.GetStub().PutState(key, data)
}

func (pl *PermissionList) GetPermission(id string) (*Permission, error) {
    key, err := pl.ctx.GetStub().CreateCompositeKey(pl.Name, []string{id})
    if err != nil {
        return nil, err
    }
    data, err := pl.ctx.GetStub().GetState(key)
    if err != nil {
        return nil, err
    }
    permission := new(Permission)
    err = json.Unmarshal(data, permission)
    return permission, err
}

func (pl *PermissionList) DelPermission(id string) error {
    key, err := pl.ctx.GetStub().CreateCompositeKey(pl.Name, []string{id})
    if err != nil {
        return err
    }
    return pl.ctx.GetStub().DelState(key)
}
