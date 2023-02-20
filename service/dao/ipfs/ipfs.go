package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"zengzhicheng/Decentralized-social-platforms/models"

	"go.uber.org/zap"

	shell "github.com/ipfs/go-ipfs-api"
)

// 定义全局变量sh
var sh *shell.Shell

// Init 初始化IPFS系统函数
func Init() (err error) {
	sh = shell.NewShell("localhost:5001")
	return err
}

// UploadIPFS 数据上传到ipfs
func UploadIPFS(str string) string {
	hash, err := sh.Add(bytes.NewBufferString(str))
	if err != nil {
		zap.L().Error("上传ipfs时错误", zap.Error(err))
	}
	return hash
}

// CatIPFS 从ipfs下载数据
func CatIPFS(hash string) string {
	read, err := sh.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(read)

	return string(body)
}

// MarshalStruct 通道序列化
func MarshalStruct(transaction models.PostIPFS) []byte {

	data, err := json.Marshal(&transaction)
	if err != nil {
		zap.L().Error("序列化err=", zap.Error(err))
	}
	return data
}

// UnmarshalStruct 数据反序列化为通道
func UnmarshalStruct(str []byte) models.PostIPFS {
	var transaction models.PostIPFS
	err := json.Unmarshal(str, &transaction)
	if err != nil {
		zap.L().Error("unmarshal err=", zap.Error(err))
	}
	return transaction
}
