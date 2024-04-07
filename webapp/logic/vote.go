package logic

import "webapp/dao/mysql"

func CreatVote(postID, userID int64) (err error) {
	return mysql.CreateVote(postID, userID)
}
