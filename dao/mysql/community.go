package mysql

import (
	"bluebell/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (data []*models.Community, err error) {
	data = make([]*models.Community, 0)
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&data, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
		zap.L().Error("query data failed", zap.Error(err))
	}
	return
}

func GetCommunityDetail(id int64) (data *models.CommunityDetail, err error) {
	data = new(models.CommunityDetail)
	sqlStr := "select community_id, community_name,introduction,create_time from community where community_id= ? "
	err = db.Get(data, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = ErrorInvalidID
		}
		zap.L().Error("query data failed", zap.Error(err))
	}
	return
}
