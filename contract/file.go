package contract

type File struct {
    // 文件ID
    FileName string `json:"file_name"`
    // 操作该文件的所需权限
    Permission string `json:"permission"`
}

