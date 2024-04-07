package mysql

func CreateVote(postID, userID int64) (err error) {
	sqlSelectVoteStr := "select count(*) from vote where post_id = ? and user_id = ? "
	var count int
	if err := db.Get(&count, sqlSelectVoteStr, postID, userID); err == nil {
		if count == 0 {
			sqlInsertVoteStr := "insert into vote (post_id,user_id) values (?,?)"
			if _, err := db.Exec(sqlInsertVoteStr, postID, userID); err != nil {
				return err
			}
			return nil
		} else {
			sqlDeleteVoteStr := "delete from vote where post_id = ? and user_id = ?"
			if _, err := db.Exec(sqlDeleteVoteStr, postID, userID); err != nil {
				return err
			}
			return nil
		}
	}
	return err
}
