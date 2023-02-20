package models

// User 用户数据
type User struct {
	UserName   string `json:"username" db:"user_name" `
	Password   string `json:"password" db:"user_password" `
	PhoneNum   string `json:"phone_num" db:"phone_num"`
	Emil       string `json:"emil" db:"emil"`
	Gender     string `json:"gender" db:"gender"`
	HeadImg    string `json:"head_img" db:"head_img"`
	CreateTime string `json:"create_time" db:"create_time"`
	UserID     int64  `json:"user_id" db:"user_id"`
	InviteId   int64  `json:"invite_id" db:"invite_id"`
	TimeStamp  int64  `json:"time_stamp" db:"time_stamp"`
}

// ApiUser 登录返回的数据
type ApiUser struct {
	Token string        `json:"token"`
	*User `json:"user"` // 嵌入用户数据

}
