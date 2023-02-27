package controller

import (
	"strconv"
	"zengzhicheng/Decentralized-social-platforms/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ----跟家相关的 ----

func FamilyHandler(c *gin.Context) {
	// 查询到所有的家（ family_id, family_name) 以列表的形式返回
	data, err := logic.GetFamilyList()
	if err != nil {
		zap.L().Error("logic.FamilyList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

// FamilyDetailHandler 家详情
func FamilyDetailHandler(c *gin.Context) {
	// 1.获取社区id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}

	// 2.根据id获取家详情
	data, err := logic.GetFamilyDetail(id)
	if err != nil {
		zap.L().Error("logic.FamilyList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

// MyFamilyHandler 属于用户的家
func MyFamilyHandler(c *gin.Context) {
	// 1.获取用户id
	idStr := c.Param("user_id")
	userId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
	}

	// 2.根据用户id获取自己家列表
	data, err := logic.GetMyFamily(userId)
	if err != nil {
		zap.L().Error("logic.FamilyList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
