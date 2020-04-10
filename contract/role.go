package contract

type Role struct {
    // 角色ID
    ID      string `json:"id"`
    Permissions []*Permission `json:"permissions"`
}
