package contract

import (
	api "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RBACContextInterface interface {
	api.TransactionContextInterface
	GetUserList() UserListInterface
	GetRoleList() RoleListInterface
	GetPermissionList() PermissionListInterface
	GetFileList() FileListInterface
	RequestFile(fileName string) (bool, error)
}

// rbac模型的上下文
type RBACContext struct {
	api.TransactionContext
	// 加入用户，角色，权限，文件等资源的数据
	// 这个地方不应该涉及到对数据的具体操作，还要往下放
	// context只应该接触上下文
	userList       *UserList
	fileList       *FileList
	permissionList *PermissionList
	roleList       *RoleList
}

func (ctx *RBACContext) GetUserList() UserListInterface {
	if ctx.userList == nil {
		ctx.userList = newUserList(ctx)
	}
	return ctx.userList
}

func (ctx *RBACContext) GetRoleList() RoleListInterface {
	if ctx.roleList == nil {
		ctx.roleList = newRoleList(ctx)
	}
	return ctx.roleList
}

func (ctx *RBACContext) GetPermissionList() PermissionListInterface {
	if ctx.permissionList == nil {
		ctx.permissionList = newPermissionList(ctx)
	}
	return ctx.permissionList
}

func (ctx *RBACContext) GetFileList() FileListInterface {
	if ctx.fileList == nil {
		ctx.fileList = newFileList(ctx)
	}
	return ctx.fileList
}

func (ctx *RBACContext) checkPermission(rId string, pIdNeeded string) (bool, error) {
	// 获取文件的permission

	r, err := ctx.GetRoleList().GetRole(rId)
	if err != nil {
		return false, err
	}
	if r == nil {
		return false, nil
	}
	for _, pId := range r.Permissions {
		if pId == pIdNeeded {
			return true, nil
		}
	}

	for _, parentId := range r.Parents {
		ok, err := ctx.checkPermission(parentId, pIdNeeded)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}

	return false, nil
}

func (ctx *RBACContext) RequestFile(fileName string) (bool, error) {
	// 首先获取文件需要的权限
	f, err := ctx.GetFileList().GetFile(fileName)
	if err != nil || f == nil {
		return false, err
	}

	p, err := ctx.GetPermissionList().GetPermission(f.Permission)
	if err != nil || p == nil {
		return false, err
	}

	xid, err := ctx.GetClientIdentity().GetX509Certificate()
	if err != nil {
		return false, err
	}
	userName := xid.Subject.CommonName

	u, err := ctx.GetUserList().GetUser(userName)
	if err != nil || u == nil {
		return false, err
	}

	for _, rId := range u.Roles {
		ok, err := ctx.checkPermission(rId, p.ID)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}

	return false, nil
}

func newRBACContext() *RBACContext {
	return &RBACContext{}
}
