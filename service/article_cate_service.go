package service

import (
	"madoka/dao"
	"madoka/engine"
	"madoka/models"
)

// ArticleCateService 文章分类Service
type ArticleCateService struct {
}

// GetAllArticleClass 获取所有的文章分类
func (s ArticleCateService) GetAllArticleClass() (*[]models.Class, error) {
	articleCateDao := dao.ArticleCateDao{DBEngine: engine.GetOrmEngine()}
	classes := make([]models.Class, 0)
	err := articleCateDao.DBEngine.Find(&classes)
	if err != nil {
		return nil, err
	}

	return &classes, err
}
