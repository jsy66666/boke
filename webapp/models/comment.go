package models

type Comment struct {
	Content string `json:"content" db:"content" binging:"required"`
}

type UpadteComment struct {
	Content string `json:"content" db:"content" binging:"required"`
}
