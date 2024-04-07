package mysql

import "webapp/models"

func NewComment(comment *models.Comment, postID, userID int64) (err error) {
	sqlStr := "insert into comment (post_id,user_id,content) values (?,?,?)"
	if _, err := db.Exec(sqlStr, postID, userID, comment.Content); err != nil {

		return err
	}
	return nil
}

func UpdateComment(updateComment *models.UpadteComment, commentID int64) (err error) {
	sqlStr := "update comment set content = ? where id = ?"
	if _, err := db.Exec(sqlStr, updateComment.Content, commentID); err != nil {
		return err
	}
	return nil
}

func DeleteComment(commentID int64) (err error) {
	sqlStr := "delete from comment where id = ?"
	if _, err := db.Exec(sqlStr, commentID); err != nil {
		return err
	}
	return nil
}
