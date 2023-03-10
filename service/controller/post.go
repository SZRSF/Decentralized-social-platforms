package controller

import (
	"fmt"
	"strconv"
	"zengzhicheng/Decentralized-social-platforms/logic"
	"zengzhicheng/Decentralized-social-platforms/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 创建作品的处理函数
func CreatePostHandler(c *gin.Context) {
	// 1.获取参数及参数的校验
	p := new(models.Post)
	works := new(models.Works)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug(" c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从 c 取到当前发请求的用户的ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	works.AuthorID = userID
	// 从 c 取到作品归属家id
	familyID, err := logic.GetFamilyID(p.FamilyName)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	works.FamilyID = familyID
	works.Content = p.Content
	works.Title = p.Title
	// 2.创建帖子
	if err := logic.CreatePost(works); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3.返回响应
	ResponseSuccess(c, nil)
}

// PostImageHandler 上传照片
func PostImageHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	response, err := logic.PostImage(file)
	fmt.Println(response)
	// 3.返回响应
	ResponseSuccess(c, response)
}

// GetPostDetailHandler 获取作品详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数（从url中获取作品的id）
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 登录用户ID
	followerId, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("CtxUserIDKey failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 2. 根据id取出帖子数据 (查数据库）
	data, err := logic.GetPostById(pid, followerId)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取作品列表的处理函数
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	// 返回响应
}

// AddCollectHandler 收藏作品
func AddCollectHandler(c *gin.Context) {
	// 1. 获取参数
	works := new(models.CollectWorks)
	if err := c.ShouldBindJSON(works); err != nil {
		zap.L().Debug(" c.ShouldBindJSON(works) error", zap.Any("err", err))
		zap.L().Error("create target with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	postId := works.ID
	// 登录用户ID
	followerId, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("CtxUserIDKey failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 2. 根据id收藏文章
	err = logic.AddCollect(postId, followerId)
	if err != nil {
		zap.L().Error("logic.AddCollect(pid, followerId)failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// DeleteCollectHandler 取消收藏作品
func DeleteCollectHandler(c *gin.Context) {
	// 1.作品id
	idStr := c.Param("worksId")
	worksId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}
	// 登录用户ID
	followerId, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("CtxUserIDKey failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 2.取消收藏事件
	err = logic.DeleteCollect(followerId, worksId)
	if err != nil {
		zap.L().Error("logic.DeleteCollect failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, nil)
}
