package contract

type Role struct {
	// 角色ID
	ID          string   `json:"id"`
	Permissions []string `json:"permissions"`
	Parents     []string `json:"parents"`// 继承的角色，用于继承权限
}
