package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"zengzhicheng/Decentralized-social-platforms/models"

	"go.uber.org/zap"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

const secret = "zengzhicheng"

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
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, user.Password)
	return
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (userID int64, err error) {
	oPassword := user.Password // 用户登录的密码
	sqlStr := `select user_id,user_name,user_password from user where user_name=?`
	err = db.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return 0, ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return 0, err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return 0, ErrorInvalidPassword
	}
	// 获取用户id
	userID = user.UserID
	return userID, err
}

// GetUserById 根据id获取用户的信息
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, user_name, phone_num, emil, gender, head_img, works_count, 
       follow_count,fans_count , like_count, collect_count, joined_family, browsing_history, 
       invite_id, time_stamp, create_time 
		from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	if err != nil {
		zap.L().Error("GetUserById failed", zap.Error(err))
	}
	return
}
