// Package dao database
package dao

import (
	"madoka/engine"
	"madoka/models"
)

// ArticleDao ad
type ArticleDao struct {
	DBEngine *engine.Orm
}

// SelectArticlePage select article page
func (d *ArticleDao) SelectArticlePage(pageSize, currentPage int) (*[]models.Article, error) {
	articles := make([]models.Article, 0)
	err := d.DBEngine.Limit(pageSize, (currentPage-1)*pageSize).Find(&articles)
	return &articles, err
}

// SelectArticleByID 以文章ID获取文章
func (d *ArticleDao) SelectArticleByID(id int) (*models.Article, error) {
	article := new(models.Article)
	_, err := d.DBEngine.ID(id).Get(article)
	return article, err
}
