package mysql

import (
	"database/sql"
	"webapp/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	return communityList, err
}

func GetCommunityDetail(communityID int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	sqlStr := "select community_id,community_name,introduction from community where community_id = ?"
	if err := db.Get(communityDetail, sqlStr, communityID); err != nil {
		if err == sql.ErrNoRows {
			err = ErrInvalidID
		}
	}
	return communityDetail, err
}
