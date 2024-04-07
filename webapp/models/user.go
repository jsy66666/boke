package models

type User struct {
	UserID   int64  `db:"user_id" binging:"required"`
	Username string `db:"user_name" binging:"required"`
	Password string `db:"password" binging:"required"`
}
