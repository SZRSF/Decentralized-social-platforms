package controller

import (
	"errors"
	"strconv"
	"zengzhicheng/Decentralized-social-platforms/dao/mysql"
	"zengzhicheng/Decentralized-social-platforms/logic"
	"zengzhicheng/Decentralized-social-platforms/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param,", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) //翻译错误
		return
	}
	// 2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("Logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist) // 用户已经存在的错误
			return
		}
		ResponseError(c, CodeServerBusy) //插入数据库失败
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)

}

// LoginHandler 登录请求函数
func LoginHandler(c *gin.Context) {
	// 1.获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param,", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam) // 登录失败，参数错误
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) //翻译错误
		return
	}
	// 2.业务逻辑处理
	data, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, data)
}

// UserDetailHandler  获取用户信息请求函数
func UserDetailHandler(c *gin.Context) {
	// 1.获取社区id
	idStr := c.Param("user_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}

	// 2.根据id获取用户详情
	data, err := logic.GetUserDetail(id)
	if err != nil {
		zap.L().Error("logic.GetUserDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

// AddFollowHandler 关注用户
func AddFollowHandler(c *gin.Context) {
	// 1.被关注者的id
	target := new(models.FollowUser)
	if err := c.ShouldBindJSON(target); err != nil {
		zap.L().Debug(" c.ShouldBindJSON(target) error", zap.Any("err", err))
		zap.L().Error("create target with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	followingId := target.Target
	// 登录用户ID
	followerId, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("CtxUserIDKey failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 2.关注用户事件
	data, err := logic.AddFollow(followerId, followingId)
	if err != nil {
		zap.L().Error(" logic.AddFollow", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

// DeleteFollowHandler 取消关注用户
func DeleteFollowHandler(c *gin.Context) {
	// 1.被关注者的id
	idStr := c.Param("target")
	followingId, err := strconv.ParseInt(idStr, 10, 64)
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
	// 2.取消关注用户事件
	data, err := logic.DeleteFollow(followerId, followingId)
	if err != nil {
		zap.L().Error("logic.DeleteFollow", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
