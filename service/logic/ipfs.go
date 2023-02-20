package logic

import (
	"zengzhicheng/Decentralized-social-platforms/dao/ipfs"
	"zengzhicheng/Decentralized-social-platforms/models"
)

// 将作品内容上传到ipfs获取hash值
func postContent(ctx string) (hash string) {
	post := models.PostIPFS{
		Content: ctx,
	}
	//结构体序列化
	data := ipfs.MarshalStruct(post)
	//上传到ipfs
	hash = ipfs.UploadIPFS(string(data))
	return hash
}

// 更具hash值从ipfs获取内容
func getContent(hash string) (post string) {
	//从ipfs下载数据
	str := ipfs.CatIPFS(hash)
	//数据反序列化
	post = ipfs.UnmarshalStruct([]byte(str)).Content
	return post
}
