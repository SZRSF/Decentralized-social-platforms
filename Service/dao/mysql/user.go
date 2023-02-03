package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"zengzhicheng/Decentralized-social-platforms/models"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

const secret = "zengzhicheng"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)

// CheckUserExist 检查指定用户名的用户是否纯在s
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where user_name = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlStr := `insert into user(user_id,user_name,user_password) value (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password // 用户登录的密码
	sqlStr := `select user_id,user_name,user_password from user where user_name=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
