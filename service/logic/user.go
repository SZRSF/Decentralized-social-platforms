package logic

import (
	"zengzhicheng/Decentralized-social-platforms/dao/mysql"
	"zengzhicheng/Decentralized-social-platforms/models"
	"zengzhicheng/Decentralized-social-platforms/pkg/jwt"
	"zengzhicheng/Decentralized-social-platforms/pkg/snowflake"
)

// 存放业务逻辑的代码

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