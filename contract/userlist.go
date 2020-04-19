package contract

import (
	api "github.com/hyperledger/fabric-contract-api-go/contractapi"
	"encoding/json"
)

type UserListInterface interface {
	AddUser(user *User) error
	GetUser(userId string) (*User, error)
	DelUser(userId string) error
}

type UserList struct {
	ctx  api.TransactionContextInterface
	Name string
}

func newUserList(ctx api.TransactionContextInterface) *UserList {
	lst := &UserList{
		ctx:  ctx,
		Name: "org.rammiah.userlist",
	}
	return lst
}

func (ul *UserList) AddUser(user *User) error {
	if len(user.Name) == 0 {
		return errBadUser
	}
	key, err := ul.ctx.GetStub().CreateCompositeKey(ul.Name, []string{user.Name})
	if err != nil {
		return err
	}
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return ul.ctx.GetStub().PutState(key, data)
}

func (ul *UserList) GetUser(name string) (*User, error) {
	key, err := ul.ctx.GetStub().CreateCompositeKey(ul.Name, []string{name})
	if err != nil {
		return nil, err
	}
	data, err := ul.ctx.GetStub().GetState(key)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	user := new(User)
	err = json.Unmarshal(data, user)
	return user, err
}

func (ul *UserList) DelUser(name string) error {
	key, err := ul.ctx.GetStub().CreateCompositeKey(ul.Name, []string{name})
	if err != nil {
		return err
	}
	return ul.ctx.GetStub().DelState(key)
}
