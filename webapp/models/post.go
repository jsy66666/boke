package models

type Post struct {
	ID          int64  `json:"id" db:"post_id"`
	Title       string `json:"title" db:"title" binging:"required"`
	Content     string `json:"content" db:"content" binging:"required"`
	AuthorID    int64  `json:"author_id" db:"author_id"`
	CommunityID int64  `json:"community_id" db:"community_id" binging:"required"`
	Status      int32  `json:"status" db:"status"`
}
