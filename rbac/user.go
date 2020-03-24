package rbac

import (
	"encoding/json"
)

// 用户信息需要和json字符串相互转换，
type User struct {
	UserName string `json:"user_name"`	
}

func (u *User) ToJson() string {
	if u == nil {
		panic("user to json cannot be nil")
	}
	// 序列化一般不会失败吧
	buf, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	return buf
}

func UserFromJson(userStr string) (*User, error) {
	
}