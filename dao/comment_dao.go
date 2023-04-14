package dao

import (
	"madoka/engine"
	"madoka/models"
)

// CommentDao 评论Dao
type CommentDao struct {
	DBEngine *engine.Orm
}

// SelectRecentComments 查询最近的评论
func (d *CommentDao) SelectRecentComments(limit int) (*[]models.CommentList, error) {
	comments := make([]models.CommentList, 0)
	err := d.DBEngine.Table("comment").Join("INNER", "article", "comment.article_id = article.id").
		Join("INNER", "user", "comment.user_id = user.id").Limit(limit).Find(&comments)
	return &comments, err
}
