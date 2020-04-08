package contract

type Role struct {
    // 角色ID
    RoleID      string `json:"role_id"`
    Permissions []*Permission
}
