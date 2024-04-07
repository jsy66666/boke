package models

type Vote struct {
	UserID int64 `db:"user_id"`
	PostID int64 `db:"post_id"`
}
