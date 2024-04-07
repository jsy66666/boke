package logic

import (
	"webapp/dao/mysql"
	"webapp/models"
	"webapp/pkg/snowflake"
)

func CreatePost(p *models.Post) error {
	//雪花算法生成postID
	var postID int64
	node, err := snowflake.NewWorker(1)
	if err != nil {
		return err
	} else {
		postID = node.GetId()
	}
	p.ID = postID
	return mysql.CreatePost(p)
}

func GetPostDetail(pid int64) (postdetail *models.Post, err error) {
	return mysql.GetPostDetail(pid)
}
