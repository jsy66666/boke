package logic

import (
	"webapp/dao/mysql"
	"webapp/models"
)

func NewComment(comment *models.Comment, postID, userID int64) (err error) {
	return mysql.NewComment(comment, postID, userID)
}

func UpdateComment(updateComment *models.UpadteComment, commentID int64) (err error) {
	return mysql.UpdateComment(updateComment, commentID)
}

func DeleteComment(commentID int64) (err error) {
	return mysql.DeleteComment(commentID)
}
