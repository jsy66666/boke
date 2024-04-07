package logic

import (
	"webapp/dao/mysql"
	"webapp/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(communityID int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetail(communityID)
}
