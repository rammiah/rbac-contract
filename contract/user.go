package contract

// 用户信息需要和json字符串相互转换，
type User struct {
    Name string `json:"name"`
    Roles []*Role `json:"roles"`
}
