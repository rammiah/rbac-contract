package contract

import (
	"encoding/json"
	api "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type FileListInterface interface {
	// 添加文件资源，并设置访问此文件需要的权限
	AddFile(file *File) error
	GetFile(fileName string) (*File, error)
	DelFile(fileName string) error
}

type FileList struct {
	ctx  api.TransactionContextInterface
	Name string
}

func newFileList(ctx api.TransactionContextInterface) *FileList {
	lst := &FileList{
		ctx:  ctx,
		Name: "org.rammiah.filelist",
	}
	return lst
}

func (fl *FileList) AddFile(file *File) error {
	key, err := fl.ctx.GetStub().CreateCompositeKey(fl.Name, []string{file.Name})
	if err != nil {
		return err
	}

	buf, _ := json.Marshal(file)
	return fl.ctx.GetStub().PutState(key, buf)
}

func (fl *FileList) GetFile(name string) (*File, error) {
	key, err := fl.ctx.GetStub().CreateCompositeKey(fl.Name, []string{name})
	if err != nil {
		return nil, err
	}
	data, err := fl.ctx.GetStub().GetState(key)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		// 文件内容为空
		return nil, nil
	}
	//fmt.Println("get old file data: ", string(data))
	f := new(File)
	err = json.Unmarshal(data, f)
	return f, err
}

func (fl *FileList) DelFile(name string) error {
	//panic("implement me")
	key, err := fl.ctx.GetStub().CreateCompositeKey(fl.Name, []string{name})
	if err != nil {
		return err
	}
	return fl.ctx.GetStub().DelState(key)
}
