package models

import "time"

type Post struct {
	FamilyName string `json:"familyName" db:"family_name" binding:"required"`
	Title      string `json:"title" db:"title" binding:"required"`
	Content    string `json:"content" db:"content" binding:"required"`
}

type CollectWorks struct {
	ID int64 `json:"worksId,string" db:"post_id"`
}

type Works struct {
	ID         int64     `json:"works_id" db:"post_id"`
	AuthorID   int64     `json:"author_id" db:"author_id"`
	FamilyName string    `json:"familyName" db:"family_name" binding:"required"`
	FamilyID   int64     `json:"family_id" db:"family_id" binding:"required"`
	Status     int32     `json:"status" db:"status"`
	Title      string    `json:"title" db:"title" binding:"required"`
	Content    string    `json:"content" db:"content" binding:"required"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

type PostIPFS struct {
	Content string `json:"content" binding:"required"`
}

// ApiPostDetail 作品详情接口
type ApiPostDetail struct {
	IsCollected   bool            `json:"is_collected"`
	IsFollowed    int16           `json:"is_followed"` // 登录用户是否关注作者
	*User         `json:"user"`   // 嵌入作者信息
	*Works                        // 嵌入帖子结构体
	*FamilyDetail `json:"family"` // 嵌入社区信息
}
