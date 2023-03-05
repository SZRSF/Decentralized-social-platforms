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

type ImageUploadResponse struct {
	Src string `json:"location"`
}

type ArticleSaveRequest struct {
	Content string `json:"content"`
}

type ArticleSaveResponse struct {
	Hash string `json:"hash"`
}

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

func AddFile(content string) (string, error) {
	// 将文章内容转换为字节数组
	data := []byte(content)
	// 将字节数组作为文件添加到 IPFS 中
	cid, err := sh.Add(bytes.NewReader(data))
	if err != nil {
		zap.L().Error("上传ipfs时错误", zap.Error(err))
		return "", err
	}
	return cid, nil
}

func PostImage(data []byte) (response ImageUploadResponse) {
	hash, err := sh.Add(bytes.NewReader(data))
	if err != nil {
		zap.L().Error("图片上传ipfs时错误", zap.Error(err))
		return
	}
	fmt.Println(hash)
	// 返回上传后的图片URL
	response = ImageUploadResponse{
		Src: "https://ipfs.io/ipfs/" + hash,
		//Src: "https://gss0.bdstatic.com/6LZ1dD3d1sgCo2Kml5_Y_D3/sys/portrait/item/tb.1.6c7e99dc.naI5KHZoNvYT7ciplr4dxw?t=1668136071",
	}
	return response
}

// CatIPFS 从ipfs下载数据
func CatIPFS(hash string) (string, error) {
	read, err := sh.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(read)

	return string(body), err
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
