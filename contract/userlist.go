package contract

import api "github.com/hyperledger/fabric-contract-api-go/contractapi"

type UserListInterface interface {
    AddUser(user *User) error
    GetUser(userId string) (*User, error)
    DelUser(userId string) error
}

type UserList struct {
    ctx api.TransactionContextInterface
}

func newUserList(ctx api.TransactionContextInterface) *UserList {
    lst := &UserList{ctx: ctx}
    return lst
}

func (ul *UserList) AddUser(user *User) error {
    panic("implement me")
}

func (ul *UserList) GetUser(userId string) (*User, error) {
    panic("implement me")
}

func (ul *UserList) DelUser(userId string) error {
    panic("implement me")
}
