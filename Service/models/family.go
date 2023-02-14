package models

import "time"

type Family struct {
	ID   int64  `json:"id" db:"family_id"`
	Name string `json:"name" db:"family_name"`
}

type FamilyDetail struct {
	ID           int64     `json:"id" db:"family_id"`
	Name         string    `json:"name" db:"family_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"createTime" db:"create_time"`
}
