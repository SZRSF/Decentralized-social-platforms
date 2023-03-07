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

// GetIsFollowed 查询登录用户是否关注文章作者
func GetIsFollowed(followerId, followingId int64) (isFollowed int16, err error) {
	sqlStr1 := `SELECT status
		FROM user_follow
		WHERE follower_id = ? AND following_id = ?;`
	err = db.Get(&isFollowed, sqlStr1, followerId, followingId)
	if err != nil {
		if err == sql.ErrNoRows {
			// 没有数据被找到
			return -1, nil
		}
		zap.L().Error("SELECT status failed", zap.Error(err))
		return 0, err
	}
	return isFollowed, err
}

// AddFollow 用户关注
func AddFollow(followerID, followingID int64) (target int16, err error) {
	// 判断是否已经存在关注关系
	isFollowed, err := GetIsFollowed(followerID, followingID)
	if err != nil {
		return isFollowed, err
	}
	if isFollowed == 1 || isFollowed == 2 {
		// 已关注或已互相关注，直接返回
		return isFollowed, nil
	}

	if isFollowed == -1 {
		// 不存在关注关系，插入一条记录
		sqlStr := `INSERT INTO user_follow (follower_id, following_id, status)
               VALUES (?, ?, 1)`
		_, err = db.Exec(sqlStr, followerID, followingID)
		if err != nil {
			zap.L().Error(" Exec failed", zap.Error(err))
			return isFollowed, err
		}
		return isFollowed, nil
	}
	// status = 0，改为1
	sqlStr := `UPDATE user_follow SET status=1 WHERE follower_id=? AND following_id=?`
	_, err = db.Exec(sqlStr, followerID, followingID)
	if err != nil {
		zap.L().Error(" Exec failed", zap.Error(err))
		return isFollowed, err
	}
	return isFollowed, nil
}

// DeleteFollow 取消关注用户
func DeleteFollow(followerID, followingID int64) (target int16, err error) {

	sqlStr := `UPDATE user_follow SET status=0 WHERE follower_id=? AND following_id=?`
	_, err = db.Exec(sqlStr, followerID, followingID)
	if err != nil {
		return target, err
	}
	target = 0
	return target, nil
}
