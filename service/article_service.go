// Package service services
package service

import (
	"madoka/dao"
	"madoka/engine"
	"madoka/models"
)

type ArticleService struct {
}

func (s *ArticleService) ListArticles(pageSize, currentPage int) (interface{}, error) {
	articleDao := dao.ArticleDao{DBEngine: engine.GetOrmEngine()}
	articlePage, err := articleDao.SelectArticlePage(pageSize, currentPage)
	return models.ArticlePageResponse{
		Pagination: models.Pagination{},
		List:       articlePage,
	}, err
}
