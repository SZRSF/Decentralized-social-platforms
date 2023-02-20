package logic

import (
	"zengzhicheng/Decentralized-social-platforms/dao/mysql"
	"zengzhicheng/Decentralized-social-platforms/models"
)

func GetFamilyList() ([]*models.Family, error) {
	// 查找数据库 查找到所有的 Family 并返回
	return mysql.GetFamilyList()
}

func GetFamilyDetail(id int64) (*models.FamilyDetail, error) {
	return mysql.GetFamilyDetailByID(id)
}
