// Package service services
package service

import (
	"madoka/dao"
	"madoka/engine"
	"madoka/models"
)

// ArticleService 文章Service
type ArticleService struct {
}

// ListArticles 文章分页
func (s *ArticleService) ListArticles(pageSize, currentPage int) (*models.ArticlePageResponse, error) {
	articleDao := dao.ArticleDao{DBEngine: engine.GetOrmEngine()}
	articlePage, err := articleDao.SelectArticlePage(pageSize, currentPage)
	return &models.ArticlePageResponse{
		Pagination: models.Pagination{},
		List:       articlePage,
	}, err
}

// GetArticle 根据文章id获取文章
func (s *ArticleService) GetArticle(id int) (*models.Article, error) {
	articleDao := dao.ArticleDao{DBEngine: engine.GetOrmEngine()}
	article, err := articleDao.SelectArticleByID(id)
	return article, err
}
