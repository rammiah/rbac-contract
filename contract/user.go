package contract

// 用户信息需要和json字符串相互转换，
type User struct {
    UserName string `json:"user_name"`
}

//func (u *User) GetSplitKey() []string {
//    return []string{u.UserName}
//}
//
//func (u *User) Serialize() ([]byte, error) {
//    if u == nil {
//        panic("user to json cannot be nil")
//    }
//    // 序列化一般不会失败吧
//    return json.Marshal(u)
//}
//
//func Deserialize(data []byte) (*User, error) {
//    return nil, nil
//}
