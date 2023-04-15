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

// SelectCommentCountOfArticle 查询某文章的评论数量
func (d *CommentDao) SelectCommentCountOfArticle(articleID int) (int, error) {
	count, err := d.DBEngine.Where("article_id = ?", articleID).Count(new(models.Comment))
	return int(count), err
}

// SelectCommentPage 查询文章的一级评论分页
func (d *CommentDao) SelectCommentPage(pageSize, currentPage, articleID int) (*[]models.Comment, error) {
	comments := make([]models.Comment, 0)
	err := d.DBEngine.Where("article_id = ? and parent_id is null", articleID).Limit(pageSize, (currentPage-1)*pageSize).Desc("create_time").Find(&comments)
	return &comments, err
}

// SelectSubComments 获取文章的二级评论
func (d *CommentDao) SelectSubComments(articleID int, parentIDList []int) (*[]models.Comment, error) {
	comments := make([]models.Comment, 0)
	err := d.DBEngine.Where("article_id = ?", articleID).In("parent_id", parentIDList).Find(&comments)
	return &comments, err
}
