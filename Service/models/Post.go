package models

import "time"

type Post struct {
	ID         int64     `json:"id" db:"post_id"`
	AuthorID   int64     `json:"author_id" db:"author_id"`
	FamilyID   int64     `json:"family_id" db:"family_id" binding:"required"`
	Status     int32     `json:"status" db:"status"`
	Title      string    `json:"title" db:"title" binding:"required"`
	Content    string    `json:"content" db:"content" binding:"required"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 作品详情接口
type ApiPostDetail struct {
	AuthorName    string          `json:"author_name"`
	*Post                         // 嵌入帖子结构体
	*FamilyDetail `json:"family"` // 嵌入社区信息
}
