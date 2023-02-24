package models

// User 用户数据
type User struct {
	UserName        string `json:"username" db:"user_name" `
	Password        string `json:"password" db:"user_password" `
	PhoneNum        string `json:"phone_num" db:"phone_num"`
	Emil            string `json:"emil" db:"emil"`
	Gender          string `json:"gender" db:"gender"`
	HeadImg         string `json:"head_img" db:"head_img"`
	CreateTime      string `json:"create_time" db:"create_time"`
	UserID          int64  `json:"user_id,string" db:"user_id"`
	InviteId        int64  `json:"invite_id" db:"invite_id"`
	TimeStamp       int64  `json:"time_stamp" db:"time_stamp"`
	WorksCount      int64  `json:"works_count" db:"works_count"`
	FollowCount     int64  `json:"follow_count" db:"follow_count"`
	FansCount       int64  `json:"fans_count" db:"fans_count"`
	LikeCount       int64  `json:"like_count" db:"like_count"`
	CollectCount    int64  `json:"collect_count" db:"collect_count"`
	JoinedFamily    int64  `json:"joined_family" db:"joined_family"`
	BrowsingHistory int64  `json:"browsing_history" db:"browsing_history"`
}

// ApiUser 登录返回的数据
type ApiUser struct {
	Token string        `json:"token"`
	*User `json:"user"` // 嵌入用户数据

}
