package contract

import (
    "fmt"
    api "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "encoding/json"
)

type RBACContract struct {
    api.Contract
}

// 创建一个用于应用链码的合约
func NewRBACContract() *RBACContract {
    contract := &RBACContract{}
    //contract.TransactionContextHandler = &RBACContract{}
    contract.TransactionContextHandler = newRBACContext()
    contract.Name = "org.rammiah.rbac"
    contract.Info.Version = "0.0.1"

    return contract
}

func (crt *RBACContract) Init(ctx RBACContextInterface) {
    fmt.Println("RBACContract initialized")
}

func (crt *RBACContract) AddFile(ctx RBACContextInterface, file string) error {
    // 不如通过json传输数据吧，也不用这么构造了
    f := new(File)
    err := json.Unmarshal([]byte(file), f)
    if err != nil {
        return err
    }
    return ctx.GetFileList().AddFile(f)
}

func (crt *RBACContract) GetFile(ctx RBACContextInterface, fileName string) (*File, error) {
    return ctx.GetFileList().GetFile(fileName)
}

func (crt *RBACContract) DelFile(ctx RBACContextInterface) {

}

func (crt *RBACContract) AddPermission(ctx RBACContextInterface) {

}

func (crt *RBACContract) GetPermission(ctx RBACContextInterface) {

}

func (crt *RBACContract) DelPermission(ctx RBACContextInterface) {

}

func (crt *RBACContract) AddRole(ctx RBACContextInterface) {

}

func (crt *RBACContract) GetRole(ctx RBACContextInterface) {

}

func (crt *RBACContract) DelRole(ctx RBACContextInterface) {

}

func (crt *RBACContract) AddUser(ctx RBACContextInterface, userName string) {
}

func (crt *RBACContract) GetUser(ctx RBACContextInterface) {

}

func (crt *RBACContract) DelUser(ctx RBACContextInterface) {

}
