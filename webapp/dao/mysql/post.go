package mysql

import (
	"fmt"
	"webapp/models"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := "insert into post (post_id,title,content,author_id,community_id) values (?,?,?,?,?)"
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return err
}

func GetPostDetail(pid int64) (postdetail *models.Post, err error) {
	postdetail = new(models.Post)
	sqlStr := "select post_id,title,content,author_id,community_id,status from post where post_id = ?"
	if err = db.Get(postdetail, sqlStr, pid); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return postdetail, nil
}
