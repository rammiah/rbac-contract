package contract

type Role struct {
    // 角色ID
    ID      string `json:"id"`
    Permissions []string `json:"permissions"`
}
