package dao

import (
	"madoka/engine"
	"madoka/models"
)

type ArticleDao struct {
	DBEngine *engine.Orm
}

func (d *ArticleDao) SelectArticlePage(pageSize, currentPage int) (*[]models.Article, error) {
	articles := make([]models.Article, 0)
	err := d.DBEngine.Limit(pageSize, (currentPage-1)*pageSize).Find(&articles)
	return &articles, err
}
