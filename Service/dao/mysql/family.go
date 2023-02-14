package mysql

import (
	"database/sql"
	"zengzhicheng/Decentralized-social-platforms/models"

	"go.uber.org/zap"
)

func GetFamilyList() (familyList []*models.Family, err error) {
	sqlStr := "select family_id, family_name from family"
	if err := db.Select(&familyList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no family in db")
			err = nil
		}
	}
	return
}

// GetFamilyDetailByID 根据id查询家详情
func GetFamilyDetailByID(id int64) (family *models.FamilyDetail, err error) {
	family = new(models.FamilyDetail)
	sqlStr := `select 
    			family_id, family_name,introduction, create_time 
				from family 
				where family_id = ?`
	if err := db.Get(family, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return family, err
}
