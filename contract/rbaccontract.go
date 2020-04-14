package contract

import (
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	api "github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/sirupsen/logrus"
)

type RBACContract struct {
	api.Contract
}

const (
	// 管理员的id，直接硬编码到代码中
	_ADMIN_ID = "Admin@org1.rammiah.org"
)

var (
	errNotPermitted = errors.New("the operation is not permitted")
	errDuplicated   = errors.New("item added is duplicated")
	errNotFound     = errors.New("item not found")
)

// 创建一个用于应用链码的合约
func NewRBACContract() *RBACContract {
	contract := &RBACContract{}
	//contract.TransactionContextHandler = &RBACContract{}
	contract.TransactionContextHandler = newRBACContext()
	contract.Name = "org.rammiah.rbac"
	contract.Info.Version = "0.0.1"

	return contract
}

func (crt *RBACContract) checkAdmin(id cid.ClientIdentity) bool {
	xid, err := id.GetX509Certificate()
	if err != nil || xid == nil {
		return false
	}
	return xid.Subject.CommonName == _ADMIN_ID
}

func (crt *RBACContract) checkOrg(id cid.ClientIdentity) bool {
	oid, err := id.GetMSPID()
	return err == nil && oid == "Org1MSP"
}

func (crt *RBACContract) AddFile(ctx RBACContextInterface, file string) error {
	if !crt.checkAdmin(ctx.GetClientIdentity()) {
		logrus.Errorf("AddFile error: %s", errNotPermitted)
		return errNotPermitted
	}
	// 不如通过json传输数据吧，也不用这么构造了
	f := new(File)
	err := json.Unmarshal([]byte(file), f)
	if err != nil {
		logrus.Errorf("marshal file json to File failed: %s", file, err)
		return err
	}
	// 检查旧文件冲突
	oldFile, err := ctx.GetFileList().GetFile(f.Name)
	if err != nil {
		logrus.Errorf("get old file error: %s", err)
		return err
	}

	if oldFile != nil {
		logrus.Errorf("file %s already exists", f.Name)
		return errDuplicated
	}
	err = ctx.GetFileList().AddFile(f)
	if err != nil {
		logrus.Errorf("AddFile error: %s", err)
		return err
	}
	return nil
}

func (crt *RBACContract) GetFile(ctx RBACContextInterface, fileName string) (*File, error) {
	if !crt.checkOrg(ctx.GetClientIdentity()) {
		logrus.Errorf("GetFile %s error: %s", fileName, errNotPermitted)
		return nil, errNotPermitted
	}
	// 获取文件
	f, err := ctx.GetFileList().GetFile(fileName)
	if err != nil {
		logrus.Errorf("GetFile %s error: %s", fileName, err)
		return nil, err
	}
	// 文件查找失败
	if f == nil {
		logrus.Errorf("GetFile failed, %s", errNotFound)
		return nil, errNotFound
	}
	return f, nil

}

func (crt *RBACContract) DelFile(ctx RBACContextInterface, fileName string) error {
	if !crt.checkAdmin(ctx.GetClientIdentity()) {
		logrus.Errorf("DelFile %s error: %s", fileName, errNotPermitted)
		return errNotPermitted
	}
	err := ctx.GetFileList().DelFile(fileName)
	//return ctx.GetFileList().DelFile(fileName)
	if err != nil {
		logrus.Errorf("DelFile error: %s", err)
		return err
	}
	return nil
}

func (crt *RBACContract) AddPermission(ctx RBACContextInterface, permission string) error {
	if !crt.checkAdmin(ctx.GetClientIdentity()) {
		logrus.Errorf("AddPermission error: %s", errNotPermitted)
		return errNotPermitted
	}
	// 从json中解析数据出来
	p := new(Permission)
	err := json.Unmarshal([]byte(permission), p)
	if err != nil {
		return err
	}
	oldP, err := ctx.GetPermissionList().GetPermission(p.ID)
	if err != nil {
		logrus.Errorf("get old permission error: %s", err)
		return err
	}

	if oldP != nil {
		logrus.Errorf("AddPermission error: %s", errDuplicated)
		return errDuplicated
	}
	err = ctx.GetPermissionList().AddPermission(p)
	if err != nil {
		logrus.Errorf("AddPermission error: %s", err)
		return err
	}
	return nil
}

func (crt *RBACContract) GetPermission(ctx RBACContextInterface, pId string) (*Permission, error) {
	if !crt.checkOrg(ctx.GetClientIdentity()) {
		logrus.Errorf("GetPermission error: %s", errNotPermitted)
		return nil, errNotPermitted
	}
	p, err := ctx.GetPermissionList().GetPermission(pId)
	if err != nil {
		logrus.Errorf("GetPermission error: %s", err)
		return nil, err
	}

	if p == nil {
		logrus.Errorf("GetPermission error: %s", errNotFound)
		return nil, errNotFound
	}

	return p, nil
}

func (crt *RBACContract) DelPermission(ctx RBACContextInterface, pId string) error {
	if !crt.checkAdmin(ctx.GetClientIdentity()) {
		logrus.Errorf("DelPermission error: %s", errNotPermitted)
		return errNotPermitted
	}

	err := ctx.GetPermissionList().DelPermission(pId)

	if err != nil {
		logrus.Errorf("DelPermission error: %s", err)
		return err
	}
	return nil
}

func (crt *RBACContract) AddRole(ctx RBACContextInterface, role string) error {
	if !crt.checkAdmin(ctx.GetClientIdentity()) {
		logrus.Errorf("AddRole failed: %s", errNotPermitted)
		return errNotPermitted
	}

	r := new(Role)
	err := json.Unmarshal([]byte(role), r)
	if err != nil {
		logrus.Errorf("marshal json to Role failed, error: %s", err)
		return err
	}

	oldR, err := ctx.GetRoleList().GetRole(r.ID)
	if err != nil {
		logrus.Errorf("get role failed, error: %s", err)
		return err
	}
	if oldR != nil {
		logrus.Errorf("AddRole error: %s", errDuplicated)
		return errDuplicated
	}

	err = ctx.GetRoleList().AddRole(r)
	if err != nil {
		logrus.Errorf("AddRole failed, error: %s", err)
		return err
	}
	return nil
}

func (crt *RBACContract) GetRole(ctx RBACContextInterface, rId string) (*Role, error) {
	if !crt.checkOrg(ctx.GetClientIdentity()) {
		logrus.Errorf("GetRole error: %s", errNotPermitted)
		return nil, errNotPermitted
	}
	r, err := ctx.GetRoleList().GetRole(rId)
	if err != nil {
		logrus.Errorf("GetRole error: %s", err)
		return nil, err
	}
	if r == nil {
		logrus.Errorf("GetRole error: %s", errNotFound)
		return nil, errNotFound
	}
	return r, nil
}

func (crt *RBACContract) DelRole(ctx RBACContextInterface, rId string) error {
	if !crt.checkAdmin(ctx.GetClientIdentity()) {
		logrus.Errorf("DelRole error: %s", errNotPermitted)
		return errNotPermitted
	}
	err := ctx.GetRoleList().DelRole(rId)
	if err != nil {
		logrus.Errorf("DelRole error: %s", err)
		return err
	}
	return nil
}

func (crt *RBACContract) AddUser(ctx RBACContextInterface, user string) error {
	if !crt.checkAdmin(ctx.GetClientIdentity()) {
		return errNotPermitted
	}
	u := new(User)
	err := json.Unmarshal([]byte(user), u)
	if err != nil {
		logrus.Errorf("marshal json to User failed, error: %s", err)
		return err
	}

	oldU, err := ctx.GetUserList().GetUser(u.Name)
	if err != nil {
		logrus.Errorf("get user error: %s", err)
		return err
	}
	if oldU != nil {
		logrus.Errorf("AddUser failed, error: %s", errDuplicated)
		return errDuplicated
	}

	err = ctx.GetUserList().AddUser(u)
	if err != nil {
		logrus.Errorf("AddUser error: %s", err)
		return err
	}

	return nil
}

func (crt *RBACContract) GetUser(ctx RBACContextInterface, uId string) (*User, error) {
	if !crt.checkOrg(ctx.GetClientIdentity()) {
		logrus.Errorf("GetUser error: %s", errNotPermitted)
		return nil, errNotPermitted
	}
	u, err := ctx.GetUserList().GetUser(uId)
	if err != nil {
		logrus.Errorf("GetUser error: %s", err)
		return nil, err
	}
	if u == nil {
		logrus.Errorf("GetUser error: %s", errNotFound)
		return nil, errNotFound
	}
	return u, nil
}

func (crt *RBACContract) DelUser(ctx RBACContextInterface, uId string) error {
	if !crt.checkAdmin(ctx.GetClientIdentity()) {
		logrus.Errorf("DelUser error: %s", errNotPermitted)
		return errNotPermitted
	}

	err := ctx.GetUserList().DelUser(uId)
	if err != nil {
		logrus.Errorf("DelUser error: %s", err)
		return err
	}
	return nil
}

func (crt *RBACContract) RequestFile(ctx RBACContextInterface, fileName string) (bool, error) {
	if !crt.checkOrg(ctx.GetClientIdentity()) {
		logrus.Errorf("RequestFile error: %s", errNotPermitted)
		return false, errNotPermitted
	}

	return ctx.RequestFile(fileName)
}
