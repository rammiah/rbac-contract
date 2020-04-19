package contract

type File struct {
	// 文件ID
	Name string `json:"name"`
	// 操作该文件的所需权限
	ReadPermission  string `json:"read_permission"`
	WritePermission string `json:"write_permission"`
	ExecPermission  string `json:"exec_permission"`
}
