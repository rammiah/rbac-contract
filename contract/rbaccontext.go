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

func newRBACContext() *RBACContext {
    return &RBACContext{}
}
