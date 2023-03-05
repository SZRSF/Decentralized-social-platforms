package logic

import (
	"bytes"
	"mime/multipart"
	"zengzhicheng/Decentralized-social-platforms/dao/ipfs"

	"go.uber.org/zap"
)

// 将作品内容上传到ipfs获取hash值
func postContent(ctx string) (cid string) {
	//post := models.PostIPFS{
	//	Content: ctx,
	//}s
	////结构体序列化
	//data := ipfs.MarshalStruct(post)
	////上传到ipfs
	//hash = ipfs.UploadIPFS(string(data))
	// 将文章内容存储到 IPFS 中
	cid, err := ipfs.AddFile(ctx)
	if err != nil {
		zap.L().Error("addFile failed", zap.Error(err))
		return
	}
	return cid
}

func PostImage(file *multipart.FileHeader) (hash ipfs.ImageUploadResponse, err error) {
	// 将图片上传到IPFS中
	reader, err := file.Open()
	if err != nil {
		zap.L().Error("file.Open failed", zap.Error(err))
		return
	}
	defer reader.Close()

	var buf bytes.Buffer
	buf.ReadFrom(reader)
	data := buf.Bytes()

	hash = ipfs.PostImage(data)
	if err != nil {
		zap.L().Error("file.Open failed", zap.Error(err))
		return
	}

	return hash, err
}

// 更具hash值从ipfs获取内容
func getContent(hash string) (post string, err error) {
	//从ipfs下载数据
	post, err = ipfs.CatIPFS(hash)
	//数据反序列化
	//post = ipfs.UnmarshalStruct([]byte(str)).Content
	return post, err
}
