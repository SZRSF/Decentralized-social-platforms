package models

import "time"

type Family struct {
	ID      int64  `json:"id" db:"family_id"`
	Name    string `json:"name" db:"family_name"`
	HeadImg string `json:"headImg" db:"family_head_img"`
}

type FamilyDetail struct {
	ID           int64     `json:"id" db:"family_id"`
	Name         string    `json:"name" db:"family_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	HeadImg      string    `json:"headImg" db:"family_head_img"`
	CreateTime   time.Time `json:"createTime" db:"create_time"`
}
