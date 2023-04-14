package service

import (
	"madoka/dao"
	"madoka/engine"
	"madoka/models"
)

// CommentService 评论Service
type CommentService struct {
}

// GetTopCommentAndTopBrowse 最新评论与最多浏览列表
func (s *CommentService) GetTopCommentAndTopBrowse() (*models.TopCommentResponse, error) {
	comments, err := getTopComments()
	if err != nil {
		return nil, err
	}

	browseList, err := getTopBrowse()
	if err != nil {
		return nil, err
	}
	topCommentResponse := new(models.TopCommentResponse)
	topCommentResponse.CommentList = comments
	topCommentResponse.BrowseList = browseList

	return topCommentResponse, err
}

func getTopComments() (*[]models.CommentList, error) {
	commentDao := dao.CommentDao{DBEngine: engine.GetOrmEngine()}
	comments, err := commentDao.SelectRecentComments(10)
	for i := range *comments {
		(*comments)[i].IsArticle = true
	}
	return comments, err
}

func getTopBrowse() (*[]models.BrowseList, error) {
	articleDao := dao.ArticleDao{DBEngine: engine.GetOrmEngine()}
	return articleDao.SelectMostBrowseArticle(10)
}
