package logic

import (
	"zengzhicheng/Decentralized-social-platforms/dao/mysql"
	"zengzhicheng/Decentralized-social-platforms/models"
	"zengzhicheng/Decentralized-social-platforms/pkg/jwt"
	"zengzhicheng/Decentralized-social-platforms/pkg/snowflake"
)

// 存放业务逻辑的代码

// SignUp 注册
func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错
		return err
	}
	// 2.生成UID
	userID := snowflake.GenID()
	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		UserName: p.Username,
		Password: p.Password,
	}
	// 3.保存进数据库
	return mysql.InsertUser(user)
}

// Login 登录
func Login(p *models.ParamLogin) (data *models.ApiUser, err error) {
	var userID int64
	user := &models.User{
		UserName: p.Username,
		Password: p.Password,
	}
	// 传递的是指针, 就能拿到user.UserID
	if userID, err = mysql.Login(user); err != nil {
		return
	}
	// 获取用户信息
	users, err := mysql.GetUserById(userID)
	if err != nil {
		return
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return
	}
	// 拼接返回数据
	data = &models.ApiUser{
		Token: token,
		User:  users,
	}
	return data, err
}

// GetUserDetail 根据Id获取用户信息
func GetUserDetail(id int64) (user *models.User, err error) {
	user, err = mysql.GetUserById(id)
	return user, err
}

// AddFollow 关注用户逻辑函数
func AddFollow(followerId, followingId int64) (target int16, err error) {
	target, err = mysql.AddFollow(followerId, followingId)
	return target, err
}

// DeleteFollow 取消关注
func DeleteFollow(followerId, followingId int64) (target int16, err error) {
	target, err = mysql.DeleteFollow(followerId, followingId)
	return target, err
}
